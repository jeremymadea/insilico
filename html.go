package main

var rootHtml = `
{{define "td" -}}
<td><a href="/?rd={{.}}"><img src="/capng?rd={{.}}&w=100&h=100" /></a></td>
{{- end}}
<html>
<head>
<title>InSilico Explorer</title>
<style type="text/css">
 body {
   font-family: sans-serif; 
   font-size: 10pt;
 }
 table { 
   margin-left: auto; 
   margin-right: auto;
 }
 img { 
   border: 1px solid white;
 }
 table, td, tr {
   padding: 0px;
   border: 0px;
   border-spacing: 0px;
 }
 #shutdown-link {
   position: absolute;
   bottom: 5px;
   right: 5px;
 }
 #random-link {
   position: absolute;
   top: 5px;
   right: 5px;
 }
 #rh-label, #rh {
   font-size: 18pt;
   font-weight: bold; 
 }
 #rh { 
   text-align: center;
   width: 9.5ch;
 }
</style>
<body>
<div style="text-align: center;">
<form action="" method="get">
  <label id="rh-label" for="rh">Ruleset:</label>
  <input id="rh" type="text" name="rh" maxlength="8" value="{{.C.RSHex}}" />
  <input type="submit" style="display: none;" />
</form>
</div>
<a id="random-link" href="/?rd=-1">Choose Randomly</a>
<table>
<tr>
{{range .MutationSetA}}{{template "td" .}}{{end}}
</tr><tr>
{{range .MutationSetB}}{{template "td" .}}{{end}}
<td colspan="4" rowspan="4" style="text-align:center;">
  <a href="/scope?rd={{.C.Ruleset}}&w=640&h=640"><img src="/capng?rd={{.C.Ruleset}}" /></a>
</td>
{{range .MutationSetC}}{{template "td" .}}{{end}}
</tr><tr>
{{range .MutationSetD}}{{template "td" .}}{{end}}
</tr><tr>
{{range .MutationSetE}}{{template "td" .}}{{end}}
<td>
</tr><tr>
{{range .MutationSetF}}{{template "td" .}}{{end}}
</tr><tr>
{{range .MutationSetG}}{{template "td" .}}{{end}}
</tr>
</table>
<a id="shutdown-link" href="/shutdown">Shutdown InSilico Explorer</a>
</body>
</html>
`

var scopeHtml = `<html>
<head>
 <title>Scope - Ruleset: {{.C.RSHex}}</title>
 <style type="text/css">
 body { 
   font-family: sans-serif; 
   font-size: 10pt;
 }
 .form-fieldset {
   display: table;
 }

 .form-field {
   display: table-row;
 }

 .field-label,
 .form-field-input,
 .form-field-comment {
   display: table-cell;
   padding: 3px 10px;
 }

 #chooser-link { 
   position: absolute;
   top: 5px; 
   right: 5px;
 }

 #shutdown-link {
   position: absolute;
   bottom: 5px;
   right: 5px;
 }

 #scope-image {
   float: left;
   padding-right: 10px;
 }
 </style>
</head>
<body>

<div id="scope-image">
<a href=""><img src="/capng{{.QueryString}}"></a>
</div>

<div id="scope-right">
<h1>RuleSet: {{.C.RSHex}}</h1>
<hr>
<div id="scope-form">
 <form action="" method="get">
 <input type="hidden" name="rd" value="{{.C.Ruleset}}">
  <fieldset><legend>Dimensions</legend>
  <div class="form-fieldset">
   <div class="form-field">
    <label class="field-label" for="w">Width:</label>
    <div class="form-field-input">
     <input id="w" type="text" size="4" name="w" value="{{.C.Width}}" />
    </div>
    <div class="form-field-comment">The width of the image.</div>
   </div>
   <div class="form-field">
    <label class="field-label" for="h">Height:</label>
    <div class="form-field-input">
     <input id="h" type="text" size="4" name="h" value="{{.C.Height}}" />
    </div>
    <div class="form-field-comment">The height of the image.</div>
   </div>
  </div>
  </fieldset>

  <fieldset><legend>Cell Colors</legend>
  <div class="form-fieldset">
   <div class="form-field">
    <label class="field-label" for="dc">Dead Cell Color:</label>
    <div class="form-field-input">
     <input id="dc" name="dc" type="color" value="#{{.C.DColorHex}}" />
    </div>
    <div class="form-field-comment"></div>
   </div>
   <div class="form-field">
    <label class="field-label" for="lc">Live Cell Color:</label>
    <div class="form-field-input">
     <input id="lc" name="lc" type="color" value="#{{.C.LColorHex}}" />
    </div>
    <div class="form-field-comment"></div>
   </div>
  </div>
  </fieldset>

  <fieldset><legend>Initialization Parameters</legend>
  <div class="form-fieldset">

  <div class="form-field">
   <label class="field-label" for="p">Initial Live %:</label>
   <div class="form-input">
     <input type="text" size="4" name="p" value="{{.C.Percentage}}" />
   </div>
   <div class="form-field-comment">Random mode only.</div>
  </div>

  <div class="form-field">
   <label class="field-label" for="m">Initialization Mode:</label>
   <div class="form-input">
     <select id="mode-drop-down" name="m">
{{.ModeOpts}}
     </select>
     <!-- <input type="text" size="6" name="m" value="{{.C.InitMode}}" /> -->
   </div> 
   <div class="form-field-comment"></div> 
  </div>  

  <div class="form-field">
   <label class="field-label" for="s">Pattern String:</label>
   <div class="form-input">
    <input type="text" size="20" name="s" value="{{.C.Pattern}}">
   </div>
   <div class="form-field-comment"></div>
  </div>
  </div>
  </fieldset>
  <br />
  <input type="submit">
 </form>
</div><!-- scope-form -->
<a id="chooser-link" href="/?rd={{.C.Ruleset}}">Back to Chooser</a>
</div>
<div id="footer">
<a id="shutdown-link" href="/shutdown">Shutdown InSilico Explorer</a>
</div>
</body>
</html>
`
