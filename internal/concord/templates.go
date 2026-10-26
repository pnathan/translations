package concord

import "github.com/pnathan/translations/internal/render"

const css = render.BaseCSS + `{{define "concord-css"}}<style>
.letters{display:flex;flex-wrap:wrap;gap:4px;direction:rtl}
.letters a{font:600 18px/1 "Noto Serif Hebrew",serif;padding:6px 9px;border:1px solid var(--line);border-radius:6px;text-decoration:none;color:var(--fg)}
.letters a.on{background:var(--accent-bg);border-color:var(--accent)}
.root h2{font:600 30px/1.3 "Noto Serif Hebrew",serif;margin:0;direction:rtl;text-align:right}
.root h2 small{font:500 13px Inter,sans-serif;color:var(--muted)}
.lemma{border-top:1px solid var(--line);margin-top:10px;padding-top:8px}
.lemma .hw{font:600 22px/1.4 "Noto Serif Hebrew",serif}
.sense{margin:6px 0 6px 0}
.sense code{font-size:12px}
.occ{display:flex;flex-wrap:wrap;gap:4px 10px;font-size:13px;margin:4px 0 0 0}
.occ a{text-decoration:none;white-space:nowrap}
.occ a .f{font:400 16px "Noto Serif Hebrew",serif}
.occ a .g{color:var(--muted)}
.note{color:var(--muted);font-size:13px}
table.ix{border-collapse:collapse;width:100%;font-size:14px}
table.ix td,table.ix th{border-top:1px solid var(--line);padding:5px 6px;text-align:left;vertical-align:top}
table.ix td.he{font-size:19px}
#q{font:16px Inter,sans-serif;padding:7px 10px;border:1px solid var(--line);border-radius:8px;background:var(--card);color:var(--fg);width:min(420px,100%)}
</style>{{end}}`

const letterHTML = `<!doctype html>
<html lang="en">
<head>{{template "head"}}{{template "concord-css"}}
<title>Concordance: {{.L.Letter}}</title>
</head>
<body>
<header class="top"><div class="wrap">
<a class="back" href="index.html">← Concordance</a>
<h1>Roots in <span class="he" lang="he">{{.L.Letter}}</span></h1>
<nav class="letters" lang="he">{{$cur := .L.Slug}}{{range .Letters}}<a href="{{.Slug}}.html"{{if eq .Slug $cur}} class="on"{{end}}>{{.Letter}}</a>{{end}}</nav>
</div></header>
<main class="wrap">
<p class="intro">Each root with its lemmas, each lemma with the lexicon's senses, and under each sense every word of the texts whose reviewed translation chose it. Select a reference to open the word in its verse.</p>
{{range .L.Roots}}
<section class="seg root" id="r{{.Index}}">
<h2 lang="he">{{if .R.Root}}{{.R.Root}} <small>root · {{.R.Count}} words</small>{{else}}{{(index .R.Lemmas 0).Headword}} <small>no root given · {{.R.Count}} words</small>{{end}}</h2>
{{if .R.Note}}<p class="note">{{.R.Note}}</p>{{end}}
<div class="note">{{.Works}}</div>
{{range .Views}}
<div class="lemma">
<span class="hw" lang="he">{{.Headword}}</span> {{if .POS}}<span class="tag">{{.POS}}</span>{{end}} {{.Gloss}} <span class="note">· {{len .Occ}} · {{.Source}} <code>{{.Key}}</code></span>
{{if .Note}}<div class="note">{{.Note}}</div>{{end}}
{{range .Groups}}
<div class="sense">{{if .Path}}<code>{{.Path}}</code> {{.Gloss}}{{else}}<span class="note">no sense chosen</span>{{end}} <span class="note">({{len .Occ}})</span>
<div class="occ">{{range .Occ}}<a href="{{link .}}"><span class="ref">{{.Ref}}</span> <span class="f" lang="he">{{.Form}}</span> <span class="g">{{.Gloss}}</span></a>{{end}}</div>
</div>
{{end}}
</div>
{{end}}
</section>
{{end}}
</main>
<script>try{const t=localStorage.getItem('theme');if(t)document.documentElement.dataset.theme=t}catch(e){}</script>
</body>
</html>`

const indexHTML = `<!doctype html>
<html lang="en">
<head>{{template "head"}}{{template "concord-css"}}
<title>Concordance</title>
</head>
<body>
<header class="top"><div class="wrap">
<a class="back" href="../index.html">← All works</a>
<h1>Concordance</h1>
<nav class="letters" lang="he">{{range .Letters}}<a href="{{.Slug}}.html">{{.Letter}}</a>{{end}}</nav>
</div></header>
<main class="wrap">
<p class="intro">{{.Groups}} roots and rootless lemmas, {{.Tokens}} words of {{range $i, $w := .Works}}{{if $i}}, {{end}}<a href="../{{$w.Slug}}/">{{$w.Title}}</a>{{end}}. Roots and senses are those of the lexica (BDB for the Bible; Jastrow and BDB, through Sefaria, for the Mishnah); the sense of each word is the one its reviewed translation chose. Built from the pipeline's data; no model writes it.</p>
<p><input id="q" type="search" placeholder="Filter: root, lemma or English gloss" aria-label="Filter"></p>
<table class="ix"><tr><th>root</th><th>lemmas</th><th>gloss</th><th>words</th></tr>
{{range .Letters}}{{$slug := .Slug}}{{range .Roots}}<tr data-k="{{.R.Root}} {{range .R.Lemmas}}{{.Headword}} {{end}}{{.Glosses}}"><td class="he" lang="he"><a href="{{$slug}}.html#r{{.Index}}">{{if .R.Root}}{{.R.Root}}{{else}}—{{end}}</a></td><td class="he" lang="he">{{range $i, $l := .R.Lemmas}}{{if $i}}, {{end}}{{$l.Headword}}{{end}}</td><td>{{.Glosses}}</td><td>{{.R.Count}}</td></tr>
{{end}}{{end}}</table>
</main>
<script>
try{const t=localStorage.getItem('theme');if(t)document.documentElement.dataset.theme=t}catch(e){}
(function(){
  const strip=s=>s.replace(/[֑-ׇ]/g,'').toLowerCase();
  const rows=[...document.querySelectorAll('tr[data-k]')].map(r=>[r,strip(r.dataset.k)]);
  document.getElementById('q').addEventListener('input',e=>{const q=strip(e.target.value.trim());rows.forEach(([r,k])=>{r.style.display=!q||k.includes(q)?'':'none'})});
})();
</script>
</body>
</html>`
