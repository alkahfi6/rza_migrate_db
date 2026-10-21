param(
    [Parameter(Mandatory=$true)][string]$SourceRoot,
    [Parameter(Mandatory=$true)][string]$ScriptDomDll,
    [Parameter(Mandatory=$true)][string]$OutputFile
)
$ErrorActionPreference = 'Stop'
Add-Type -Path $ScriptDomDll
$code = @'
using System;
using System.IO;
using System.Collections;
using System.Collections.Generic;
using System.Reflection;
using Microsoft.SqlServer.TransactSql.ScriptDom;
public static class UatAuditAst {
    static Dictionary<Type,PropertyInfo[]> cache = new Dictionary<Type,PropertyInfo[]>();
    static void Walk(TSqlFragment f, int parent, string role, List<Dictionary<string,object>> nodes) {
        if(f == null) return;
        int id=nodes.Count;
        var n=new Dictionary<string,object> {{"t",f.GetType().Name},{"p",parent},{"r",role},{"l",f.StartLine},{"o",f.StartOffset},{"z",f.FragmentLength}};
        nodes.Add(n);
        PropertyInfo[] props;
        if(!cache.TryGetValue(f.GetType(),out props)) { props=f.GetType().GetProperties(); cache[f.GetType()]=props; }
        foreach(var prop in props) {
            if(prop.GetIndexParameters().Length!=0 || prop.Name=="ScriptTokenStream" || prop.Name=="FirstTokenIndex" || prop.Name=="LastTokenIndex" || prop.Name=="StartOffset" || prop.Name=="StartLine" || prop.Name=="StartColumn" || prop.Name=="FragmentLength") continue;
            object v=prop.GetValue(f);
            if(v==null) continue;
            if(v is TSqlFragment) Walk((TSqlFragment)v,id,prop.Name,nodes);
            else if(v is string || v is bool || v.GetType().IsEnum) n[prop.Name]=v.ToString();
            else if(v is IEnumerable) foreach(var child in (IEnumerable)v) if(child is TSqlFragment) Walk((TSqlFragment)child,id,prop.Name,nodes);
        }
    }
    public static Dictionary<string,object> Parse(string path, string rel) {
        string sql=File.ReadAllText(path);
        IList<ParseError> errors;
        TSqlFragment root=new TSql170Parser(true).Parse(new StringReader(sql),out errors);
        var es=new List<Dictionary<string,object>>();
        foreach(var e in errors) es.Add(new Dictionary<string,object>{{"line",e.Line},{"number",e.Number},{"message",e.Message}});
        var ns=new List<Dictionary<string,object>>();
        Walk(root,-1,"Root",ns);
        return new Dictionary<string,object>{{"path",rel},{"errors",es},{"nodes",ns}};
    }
}
'@
Add-Type -TypeDefinition $code -ReferencedAssemblies @($ScriptDomDll, 'System.Collections', 'System.Runtime', 'netstandard')
$root = (Resolve-Path -LiteralPath $SourceRoot).Path
$writer = [IO.StreamWriter]::new($OutputFile, $false, [Text.UTF8Encoding]::new($false))
try {
    $files = Get-ChildItem -LiteralPath $root -Recurse -File -Filter '*.sql' | Sort-Object FullName
    foreach($file in $files) {
        $rel = [IO.Path]::GetRelativePath($root, $file.FullName).Replace('\','/')
        $data = [UatAuditAst]::Parse($file.FullName, $rel)
        $data['sha256'] = (Get-FileHash -LiteralPath $file.FullName -Algorithm SHA256).Hash.ToLowerInvariant()
        $writer.WriteLine([System.Text.Json.JsonSerializer]::Serialize($data, $data.GetType(), [System.Text.Json.JsonSerializerOptions]::new()))
    }
    Write-Output ('Parsed {0} source files' -f $files.Count)
} finally { $writer.Dispose() }
