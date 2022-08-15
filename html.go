/* MIT License

Copyright (c) 2022 Jeremy Madea

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
 #builder-link {
   position: absolute;
   top: 20px;
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
<a id="builder-link" href="/builder?rh={{.C.RSHex}}">Builder</a>
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

 #builder-link { 
   position: absolute;
   top: 20px; 
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
<h1>Ruleset: {{.C.RSHex}}</h1>
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
<a id="chooser-link" href="/?rd={{.C.Ruleset}}">Chooser</a>
<a id="builder-link" href="/builder?rd={{.C.Ruleset}}">Builder</a>
</div>
<div id="footer">
<a id="shutdown-link" href="/shutdown">Shutdown InSilico Explorer</a>
</div>
</body>
</html>
`

var builderHtml = `
<html>
<head>
<title>In Silico Explorer</title>
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
/* table, td, tr {
   padding: 0px;
   border: 0px;
   border-spacing: 0px;
 }*/
 td { 
   border: 1px solid black;
 }
 #shutdown-link {
   position: absolute;
   bottom: 5px;
   right: 5px;
 }
 #chooser-link { 
   position: absolute;
   top: 20px; 
   right: 5px;
 }
 #random-link {
   position: absolute;
   top: 5px;
   right: 5px;
 }
 #scope-link {
   position: absolute;
   top: 35px;
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
 .tbl-container { 
   float: left;
   border-right: 3px double black;
 }
 .first-container { 
   border-left: 3px double black;
   clear: left;
 }
 .spacer { 
   border: none;
 }
 .cell { 
   border: 1px solid black; 
 }
 .next { 
   border-color: red;
 }
 .nbr1, .nbr2 { 
   border: 1px solid black;
 }

 .blk { 
   background: black;
 }
 .wht { 
   background: white;
 }

 .ldsep { 
   /*background: lightgrey; */
   height: 3em;
 }

 .img-container { 
   float: left;
   padding-left: 30px;
 }

 td a { 
   text-decoration: none;
   display: block; 
   width: 100%;
 }

</style>
<body>

<!-- <h1>Ruleset: {{.C.RSHex}}</h1> -->
<div style="float: left; clear: right;">
<form action="" method="get">
  <label id="rh-label" for="rh">Ruleset:</label>
  <input id="rh" type="text" name="rh" maxlength="8" value="{{.C.RSHex}}" />
  <input type="submit" style="display: none;" />
</form>
</div>
<a id="random-link" href="?rd=-1">Choose Randomly</a>
<a id="chooser-link" href="/?rd={{.C.Ruleset}}">Chooser</a>
<a id="scope-link" href="/scope?rd={{.C.Ruleset}}&w=640&h=640">Scope</a>
<!-- BEGIN GENERATED CODE -->
<div class="tbl-container first-container"><table>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 0}} cell next"><a href="?rd={{index .RSNew 0}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 1}} cell next"><a href="?rd={{index .RSNew 1}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 2}} cell next"><a href="?rd={{index .RSNew 2}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 3}} cell next"><a href="?rd={{index .RSNew 3}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer ldsep" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 4}} cell next"><a href="?rd={{index .RSNew 4}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 5}} cell next"><a href="?rd={{index .RSNew 5}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 6}} cell next"><a href="?rd={{index .RSNew 6}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 7}} cell next"><a href="?rd={{index .RSNew 7}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
</table></div>
<div class="tbl-container"><table>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 8}} cell next"><a href="?rd={{index .RSNew 8}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 9}} cell next"><a href="?rd={{index .RSNew 9}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 10}} cell next"><a href="?rd={{index .RSNew 10}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 11}} cell next"><a href="?rd={{index .RSNew 11}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer ldsep" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 12}} cell next"><a href="?rd={{index .RSNew 12}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 13}} cell next"><a href="?rd={{index .RSNew 13}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 14}} cell next"><a href="?rd={{index .RSNew 14}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 15}} cell next"><a href="?rd={{index .RSNew 15}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
</table></div>
<div class="tbl-container"><table>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 16}} cell next"><a href="?rd={{index .RSNew 16}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 17}} cell next"><a href="?rd={{index .RSNew 17}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 18}} cell next"><a href="?rd={{index .RSNew 18}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 19}} cell next"><a href="?rd={{index .RSNew 19}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer ldsep" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 20}} cell next"><a href="?rd={{index .RSNew 20}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 21}} cell next"><a href="?rd={{index .RSNew 21}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 22}} cell next"><a href="?rd={{index .RSNew 22}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 23}} cell next"><a href="?rd={{index .RSNew 23}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
</table></div>
<div class="tbl-container"><table>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 24}} cell next"><a href="?rd={{index .RSNew 24}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 25}} cell next"><a href="?rd={{index .RSNew 25}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 26}} cell next"><a href="?rd={{index .RSNew 26}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 27}} cell next"><a href="?rd={{index .RSNew 27}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer ldsep" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 28}} cell next"><a href="?rd={{index .RSNew 28}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 blk">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 29}} cell next"><a href="?rd={{index .RSNew 29}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 blk">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 30}} cell next"><a href="?rd={{index .RSNew 30}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
<tr><td class="spacer" colspan="5">&nbsp;</td></tr>
<tr>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="cell wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr1 wht">&nbsp;&nbsp;&nbsp;</td>
<td class="nbr2 wht">&nbsp;&nbsp;&nbsp;</td>
</tr>
<tr>
<td class="spacer" colspan="2"></td>
<td class="{{index .Color 31}} cell next"><a href="?rd={{index .RSNew 31}}"><div>&nbsp;</div></a></td>
<td class="spacer" colspan="2"></td>
</tr>
</table></div>
<!--END GENERATED CODE -->

<div class="img-container">
<img src="/capng?rd={{.C.Ruleset}}&w=480&h=480" />
</div>

<a id="shutdown-link" href="/shutdown">Shutdown InSilico Explorer</a>
</body>
</html>
`
