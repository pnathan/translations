package render

// BaseCSS defines the template "head" (fonts and colours) shared by every page.
const BaseCSS = baseCSS

const baseCSS = `{{define "head"}}<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Cormorant+Garamond:ital,wght@0,500;0,600;1,500&family=Inter:wght@400;500;600&family=Noto+Serif+Hebrew:wght@400;600&display=swap" rel="stylesheet">
<style>
:root{--bg:#f7f4ee;--card:#fffdf9;--fg:#1d1b17;--muted:#6b645a;--line:#e3dccf;--accent:#8a4b16;--accent-bg:#f3e6d6;--hi:#fff1b8;--tag:#ece6da}
@media (prefers-color-scheme:dark){:root:not([data-theme="light"]){--bg:#141311;--card:#1c1a17;--fg:#ece6dc;--muted:#a39a8c;--line:#34302a;--accent:#e2a565;--accent-bg:#3a2a1a;--hi:#4a3f12;--tag:#2a2621}}
:root[data-theme="dark"]{--bg:#141311;--card:#1c1a17;--fg:#ece6dc;--muted:#a39a8c;--line:#34302a;--accent:#e2a565;--accent-bg:#3a2a1a;--hi:#4a3f12;--tag:#2a2621}
*{box-sizing:border-box}
html{-webkit-text-size-adjust:100%}
body{margin:0;background:var(--bg);color:var(--fg);font:16px/1.55 Inter,system-ui,sans-serif}
a{color:var(--accent)}
.wrap{max-width:1100px;margin:0 auto;padding:0 16px}
header.top{border-bottom:1px solid var(--line);background:var(--card);position:sticky;top:0;z-index:5}
header.top .wrap{display:flex;flex-wrap:wrap;gap:8px 16px;align-items:center;padding-top:10px;padding-bottom:10px}
.back{font-size:14px;text-decoration:none}
h1{font:600 26px/1.2 "Cormorant Garamond",serif;margin:0}
h1 .he,.he{font-family:"Noto Serif Hebrew","SBL Hebrew","Ezra SIL",serif}
nav.sections{display:flex;gap:6px;flex-wrap:wrap}
nav.sections a{font-size:13px;padding:3px 9px;border:1px solid var(--line);border-radius:999px;text-decoration:none;color:var(--fg)}
nav.sections a.on{background:var(--accent-bg);border-color:var(--accent)}
.toggles{display:flex;gap:6px;flex-wrap:wrap;margin-left:auto}
.toggles button{font:500 13px Inter,sans-serif;padding:4px 10px;border-radius:6px;border:1px solid var(--line);background:var(--bg);color:var(--fg);cursor:pointer}
.toggles button[aria-pressed="true"]{background:var(--accent-bg);border-color:var(--accent)}
main{padding:20px 0 120px}
.intro{color:var(--muted);font-size:14px;max-width:70ch}
.seg{background:var(--card);border:1px solid var(--line);border-radius:10px;padding:14px 16px;margin:14px 0}
.seg.para-end{margin-bottom:28px}
.ref{font:600 12px Inter,sans-serif;color:var(--muted);letter-spacing:.04em}
.heb{direction:rtl;text-align:right;font:400 26px/2.1 "Noto Serif Hebrew","SBL Hebrew","Ezra SIL",serif;margin:6px 0 10px}
.w{display:inline;cursor:pointer;border-radius:4px;padding:0 1px;background:none;border:0;color:inherit;font:inherit}
.w:hover,.w:focus-visible{background:var(--accent-bg);outline:none}
.w.sel{background:var(--hi)}
.w.anch{text-decoration:underline dotted var(--muted);text-underline-offset:6px}
.w .gl{display:none}
body.inter .heb{line-height:1.4}
body.inter .w{display:inline-flex;flex-direction:column;align-items:center;vertical-align:top;margin:0 3px 10px}
body.inter .w .gl{display:block;direction:ltr;font:400 12px/1.25 Inter,sans-serif;color:var(--muted);max-width:9em;text-align:center}
.en{font:500 20px/1.5 "Cormorant Garamond",serif}
.lit{display:none;color:var(--muted);font-size:14px;margin-top:4px}
body.literal .lit{display:block}
details{margin-top:8px;font-size:14px}
summary{cursor:pointer;color:var(--muted);font-weight:500}
details ul{margin:6px 0;padding-left:20px}
.comm{border-left:3px solid var(--line);padding:4px 0 4px 10px;margin:8px 0}
.comm.lit-on{border-color:var(--accent);background:var(--accent-bg)}
.tag{display:inline-block;font:600 11px Inter,sans-serif;padding:1px 7px;border-radius:999px;background:var(--tag);color:var(--muted);margin-left:4px}
.tag.norm{color:var(--accent)}
.hebtext{direction:rtl;text-align:right;font:400 16px/1.8 "Noto Serif Hebrew",serif}
body.nocomm details.comms{display:none}
aside#panel{position:fixed;right:16px;top:84px;width:380px;max-height:calc(100vh - 100px);overflow:auto;background:var(--card);border:1px solid var(--line);border-radius:12px;padding:14px 16px;box-shadow:0 8px 28px rgba(0,0,0,.18);display:none;z-index:10}
aside#panel.open{display:block}
#panel .pw{font:400 34px/1.6 "Noto Serif Hebrew",serif;direction:rtl;text-align:right}
#panel h3{font:600 13px Inter,sans-serif;color:var(--muted);margin:12px 0 4px;text-transform:uppercase;letter-spacing:.05em}
#panel table{border-collapse:collapse;width:100%;font-size:14px}
#panel td{border-top:1px solid var(--line);padding:4px 6px 4px 0;vertical-align:top}
#panel .senses li.on{background:var(--hi);border-radius:4px}
#panel .x{float:right;border:0;background:none;font-size:20px;cursor:pointer;color:var(--muted)}
.vocab table{border-collapse:collapse;width:100%;font-size:14px}
.vocab td,.vocab th{border-top:1px solid var(--line);padding:4px 6px;text-align:left}
footer{border-top:1px solid var(--line);color:var(--muted);font-size:13px;padding:16px 0 40px}
code{font-size:12px}
#stage{position:fixed;inset:0;z-index:20;display:none;perspective:1600px;perspective-origin:50% 40%;background:color-mix(in srgb,var(--bg) 70%,transparent)}
#stage.open{display:block}
#stage .deck{position:absolute;left:50%;top:46%;width:min(640px,calc(100vw - 32px));transform-style:preserve-3d}
#stage .card{position:absolute;left:0;top:0;width:100%;transform:translate(-50%,-50%);background:var(--card);border:1px solid var(--line);border-radius:12px;padding:14px 16px;box-shadow:0 18px 50px rgba(0,0,0,.28);transition:transform .55s cubic-bezier(.2,.7,.2,1),opacity .55s;cursor:pointer;backface-visibility:hidden}
#stage .card .heb{font-size:22px;line-height:1.9;margin:4px 0 6px}
#stage .card .hit{background:var(--hi);border-radius:4px;box-shadow:0 0 0 2px var(--accent)}
#stage .card .en{font-size:17px}
#stage .card .g{font-size:13px;color:var(--muted)}
#stage .bar{position:absolute;left:16px;right:16px;top:12px;display:flex;flex-wrap:wrap;gap:8px 14px;align-items:center;background:var(--card);border:1px solid var(--line);border-radius:10px;padding:8px 12px;z-index:2}
#stage .bar .rt{font:600 24px/1.2 "Noto Serif Hebrew",serif}
#stage .bar button{font:500 13px Inter,sans-serif;padding:4px 10px;border-radius:6px;border:1px solid var(--line);background:var(--bg);color:var(--fg);cursor:pointer}
#stage .bar .sp{margin-left:auto;display:flex;gap:6px;align-items:center}
main.wrap{transition:transform .7s cubic-bezier(.2,.7,.2,1),opacity .7s;transform-origin:0 50%}
body.swung main.wrap{transform:perspective(1400px) rotateY(58deg);opacity:.4;pointer-events:none}
@media (prefers-reduced-motion:reduce){main.wrap,#stage .card{transition:none}}
@media (max-width:820px){#stage .deck{top:60%}#stage .card{max-height:60vh;overflow:auto}#stage .card .heb{font-size:20px}}
@media (max-width:820px){aside#panel{left:0;right:0;top:auto;bottom:0;width:auto;max-height:55vh;border-radius:12px 12px 0 0}.heb{font-size:23px}.toggles{margin-left:0}}
</style>{{end}}`

const sectionHTML = `<!doctype html>
<html lang="en">
<head>{{template "head"}}
<title>{{.Cfg.Title}}: {{.Section.Name}}</title>
</head>
<body>
<header class="top"><div class="wrap">
<a class="back" href="index.html">← {{.Cfg.Title}}</a>
<h1>{{.Section.Name}} <span class="he" lang="he">{{.Section.He}}</span></h1>
<nav class="sections">{{$cur := .Section.Slug}}{{range .Sections}}<a href="{{.Slug}}.html"{{if eq .Slug $cur}} class="on"{{end}}>{{.Name}}</a>{{end}}</nav>
<div class="toggles">
<button data-t="inter" aria-pressed="false">Interlinear</button>
<button data-t="literal" aria-pressed="false">Literal</button>
<button data-t="nocomm" aria-pressed="false">Hide commentary</button>
<button data-t="three" aria-pressed="false" title="Select a word to see every verse that uses its root">3D</button>
<button id="theme">Theme</button>
</div>
</div></header>
<main class="wrap">
<p class="intro">{{.Cfg.RefName}} {{.First}}–{{.Last}}. Select a Hebrew word for its morphology, lemma, root, senses and the comments anchored to it. A dotted underline marks a word a commentator quotes.</p>
{{range .Segs}}
<section class="seg{{if .Para}} para-end{{end}}" id="s{{.Ref}}">
<div class="ref">{{.Ref}}{{if .Para}} · {{.Para}}{{end}}</div>
<div class="heb" lang="he">{{range .Words}}<button class="w{{if .Anchored}} anch{{end}}" data-id="{{.ID}}"><span>{{.Text}}{{if .Maqaf}}־{{end}}{{.Punct}}</span><span class="gl">{{.Gloss}}</span></button>{{if not .Maqaf}} {{end}}{{end}}</div>
{{if .Missing}}<p class="en"><em>Not yet translated.</em></p>{{else}}
<p class="en">{{.English}}</p>
<p class="lit">{{.Literal}}</p>
{{if .Grammar}}<details><summary>Grammar</summary><ul>{{range .Grammar}}<li>{{.}}</li>{{end}}</ul>{{if .Note}}<p>{{.Note}}</p>{{end}}</details>{{end}}
{{end}}
{{if .Comments}}<details class="comms"><summary>Commentary ({{len .Comments}})</summary>
{{range .Comments}}<div class="comm" id="c-{{.ID}}"><strong>{{.Commentator}}</strong>{{if .Normative}}<span class="tag norm">normative</span>{{end}}{{if .Relation}}<span class="tag">{{.Relation}}</span>{{end}} <span class="he" lang="he">{{.Span}}</span>
{{if .Summary}}<div>{{.Summary}}</div>{{end}}
<details><summary>Hebrew</summary><div class="hebtext" lang="he">{{.Hebrew}}</div></details></div>{{end}}
</details>{{end}}
{{if .Context}}<details><summary>Ancient Near East</summary><ul>{{range .Context}}<li>{{.}}</li>{{end}}</ul></details>{{end}}
{{if .Apparatus}}<details><summary>Apparatus</summary><ul>{{range .Apparatus}}<li>{{.}}</li>{{end}}</ul></details>{{end}}
</section>
{{end}}
<section class="seg vocab"><details><summary>Vocabulary of this section ({{len .Vocab}} lemmas, by frequency)</summary>
<table><tr><th>lemma</th><th>root</th><th>gloss</th><th>count</th></tr>
{{range .Vocab}}<tr><td class="he" lang="he">{{.Lemma}}</td><td class="he" lang="he">{{.Root}}</td><td>{{.Gloss}}</td><td>{{.Count}}</td></tr>{{end}}
</table></details></section>
</main>
<aside id="panel" aria-live="polite"></aside>
<div id="stage" role="dialog" aria-modal="true" aria-label="Verses with this root"><div class="bar"></div><div class="deck"></div></div>
<footer><div class="wrap"><ul>{{.Provenance}}</ul></div></footer>
<script>
const DATA = {{.Data}};
const WORK = {{.Cfg.Slug}};
(function(){
  const root=document.documentElement, body=document.body;
  const store={get(k){try{return localStorage.getItem(k)}catch(e){return null}},set(k,v){try{localStorage.setItem(k,v)}catch(e){}}};
  const t=store.get('theme'); if(t) root.dataset.theme=t;
  document.getElementById('theme').onclick=()=>{const d=root.dataset.theme==='dark'||(!root.dataset.theme&&matchMedia('(prefers-color-scheme: dark)').matches);root.dataset.theme=d?'light':'dark';store.set('theme',root.dataset.theme)};
  document.querySelectorAll('.toggles button[data-t]').forEach(b=>{
    const k=b.dataset.t; if(store.get('t-'+k)==='1'){body.classList.add(k);b.setAttribute('aria-pressed','true')}
    b.onclick=()=>{const on=body.classList.toggle(k);b.setAttribute('aria-pressed',on);store.set('t-'+k,on?'1':'0')};
  });
  const panel=document.getElementById('panel');
  const esc=s=>String(s==null?'':s).replace(/[&<>"]/g,c=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;'}[c]));
  let sel=null;
  function show(id){
    const d=DATA.tokens[id]; if(!d) return;
    if(sel) sel.classList.remove('sel');
    document.querySelectorAll('.comm.lit-on').forEach(x=>x.classList.remove('lit-on'));
    sel=document.querySelector('.w[data-id="'+id+'"]'); if(sel) sel.classList.add('sel');
    let h='<button class="x" aria-label="Close">×</button><div class="pw" lang="he">'+esc(d.t)+'</div>';
    if(d.g) h+='<div><strong>'+esc(d.g)+'</strong></div>';
    if(d.n) h+='<div>'+esc(d.n)+'</div>';
    if(d.m&&d.m.length){h+='<h3>Morphemes</h3><table>';d.m.forEach(m=>{h+='<tr><td class="he" lang="he">'+esc(m.f)+'</td><td>'+esc(m.l||'')+'</td><td>'+esc(m.p||'')+(m.x?': '+esc(m.x):'')+'</td></tr>'});h+='</table>'}
    if(d.r) h+='<h3>Root</h3><div class="he" lang="he">'+esc(d.r)+'</div>';
    if(d.a&&d.a.length) h+='<h3>Accent</h3><div>'+esc(d.a.join(', '))+'</div>';
    const e=d.k&&DATA.lex[d.k];
    if(e){h+='<h3>Lexicon</h3><div><span class="he" lang="he">'+esc(e.headword)+'</span> '+esc(e.pos||'')+(e.gloss?' — '+esc(e.gloss):'')+'</div><div style="font-size:12px;color:var(--muted)">'+esc(e.source)+'</div>';
      if(e.senses&&e.senses.length){h+='<ul class="senses">';e.senses.forEach(s=>{h+='<li'+(s.path===d.s?' class="on"':'')+'><code>'+esc(s.path)+'</code> '+esc(s.gloss)+'</li>'});h+='</ul>'}}
    if(d.kt) h+='<h3>Ketiv</h3><div class="he" lang="he">'+esc(d.kt)+'</div>';
    if(d.alt) h+='<h3>Other witness</h3><div class="he" lang="he">'+esc(d.alt)+'</div>';
    if(d.wn&&d.wn.length) h+='<h3>Textual notes</h3><ul>'+d.wn.map(x=>'<li>'+esc(x)+'</li>').join('')+'</ul>';
    if(d.c&&d.c.length){h+='<h3>Commentary on this word</h3><ul>';d.c.forEach(c=>{const el=document.getElementById('c-'+c);if(el){el.classList.add('lit-on');const s=el.querySelector('strong');h+='<li><a href="#c-'+esc(c)+'" data-open="'+esc(c)+'">'+esc(s?s.textContent:c)+'</a></li>'}});h+='</ul>'}
    if(d.src) h+='<div style="font-size:12px;color:var(--muted);margin-top:10px">Morphology: '+esc(d.src)+'</div>';
    panel.innerHTML=h; panel.classList.add('open');
    panel.querySelector('.x').onclick=()=>{panel.classList.remove('open');if(sel)sel.classList.remove('sel')};
    panel.querySelectorAll('[data-open]').forEach(a=>a.onclick=()=>{const el=document.getElementById('c-'+a.dataset.open);if(el){const dd=el.closest('details.comms');if(dd)dd.open=true}});
  }
  // The concordance (docs/lexicon), loaded on first use.
  let IX=null; const VERSES={};
  const getJSON=u=>fetch(u).then(r=>{if(!r.ok)throw new Error(u+': '+r.status);return r.json()});
  const index=()=>IX?Promise.resolve(IX):getJSON('../lexicon/index.json').then(x=>IX=x);
  const verses=w=>VERSES[w]?Promise.resolve(VERSES[w]):getJSON('../lexicon/verses-'+w+'.json').then(x=>VERSES[w]=x);
  const refOf=id=>id.slice(0,id.lastIndexOf('.'));
  function concordLink(id){
    index().then(ix=>{
      const i=ix.tok[WORK+':'+id]; if(i==null||!panel.classList.contains('open')||!sel||sel.dataset.id!==id) return;
      const r=ix.roots[i], d=document.createElement('div');
      d.innerHTML='<h3>Concordance</h3><a href="../lexicon/'+esc(r.p)+'"><span class="he" lang="he">'+esc(r.r||r.h[0])+'</span> · '+r.o.length+' words</a> · <button class="c3" style="font:500 12px Inter,sans-serif;border:1px solid var(--line);border-radius:6px;background:var(--bg);color:var(--fg);cursor:pointer">3D</button>';
      panel.appendChild(d); d.querySelector('.c3').onclick=()=>three(id);
    }).catch(()=>{});
  }
  // 3D: the page swings back and the verses that share the word's root
  // stand behind it as a deck of cards, the selected verse first.
  const stage=document.getElementById('stage'), deck=stage.querySelector('.deck'), bar=stage.querySelector('.bar');
  let cards=[], cur=0;
  function place(){
    cards.forEach((c,i)=>{const k=i-cur;
      if(k<-1||k>10){c.style.display='none';return}
      c.style.display='';
      c.style.transform=k<0?'translate(-50%,-50%) translate3d(-120%,0,200px) rotateY(-25deg)':'translate(-50%,-50%) translate3d('+(k*46)+'px,'+(-k*34)+'px,'+(-k*230)+'px) rotateY('+(-k*4)+'deg)';
      c.style.opacity=k<0?0:Math.max(0,1-k*.11); c.style.zIndex=100-k;
    });
    const n=bar.querySelector('.n'); if(n) n.textContent=(cur+1)+' of '+cards.length;
  }
  function closeThree(){stage.classList.remove('open');body.classList.remove('swung');deck.innerHTML='';cards=[];if(sel)sel.focus()}
  function three(id){
    index().then(ix=>{
      const i=ix.tok[WORK+':'+id]; if(i==null) return;
      const r=ix.roots[i];
      // one card per verse; this verse first, then this work, then the others
      const by=new Map();
      r.o.forEach(([w,sec,tid,g])=>{const k=w+'|'+refOf(tid);if(!by.has(k))by.set(k,{w,sec,ref:refOf(tid),ids:[],gl:[]});const v=by.get(k);v.ids.push(tid);if(g)v.gl.push(g)});
      const here=WORK+'|'+refOf(id);
      const rank=v=>(v.w+'|'+v.ref)===here?0:v.w===WORK?1:2;
      const list=[...by.values()].sort((a,b)=>rank(a)-rank(b));
      return Promise.all([...new Set(list.map(v=>v.w))].map(w=>verses(w))).then(()=>{
        bar.innerHTML='<span class="rt" lang="he">'+esc(r.r||r.h[0])+'</span><span>'+esc([...new Set(r.h)].join(', '))+'</span>'+(r.g?'<span class="g" style="color:var(--muted)">'+esc(r.g)+'</span>':'')+'<a href="../lexicon/'+esc(r.p)+'">concordance</a><span class="sp"><button class="pv" aria-label="Previous">↑</button><span class="n"></span><button class="nx" aria-label="Next">↓</button><button class="cl" aria-label="Close">×</button></span>';
        deck.innerHTML=''; cards=list.map(v=>{
          const vs=VERSES[v.w][v.ref]; const c=document.createElement('article'); c.className='card';
          const hit=new Set(v.ids);
          c.innerHTML='<div class="ref">'+esc(ix.works[v.w]?ix.works[v.w].t:v.w)+' '+esc(v.ref)+'</div><div class="heb" lang="he">'+(vs?vs.w.map(([t,x])=>'<span'+(hit.has(t)?' class="hit"':'')+'>'+esc(x)+'</span>').join(' '):'')+'</div><p class="en">'+esc(vs?vs.en:'')+'</p><div class="g">'+esc([...new Set(v.gl)].join(' · '))+'</div>';
          c.onclick=()=>{const j=cards.indexOf(c);if(j===cur)location.href='../'+v.w+'/'+v.sec+'.html#t'+v.ids[0];else{cur=j;place()}};
          deck.appendChild(c); return c;
        });
        cur=0; panel.classList.remove('open'); body.classList.add('swung'); stage.classList.add('open'); place();
        bar.querySelector('.pv').onclick=()=>{cur=Math.max(0,cur-1);place()};
        bar.querySelector('.nx').onclick=()=>{cur=Math.min(cards.length-1,cur+1);place()};
        bar.querySelector('.cl').onclick=closeThree; bar.querySelector('.nx').focus();
      });
    }).catch(e=>{panel.insertAdjacentHTML('beforeend','<div style="color:var(--muted)">The concordance could not be loaded.</div>')});
  }
  stage.addEventListener('wheel',ev=>{if(!cards.length)return;ev.preventDefault();cur=Math.max(0,Math.min(cards.length-1,cur+(ev.deltaY>0?1:-1)));place()},{passive:false});
  stage.addEventListener('click',ev=>{if(ev.target===stage)closeThree()});
  document.addEventListener('click',ev=>{const w=ev.target.closest('.w');if(!w)return;show(w.dataset.id);if(body.classList.contains('three'))three(w.dataset.id);else concordLink(w.dataset.id)});
  document.addEventListener('keydown',ev=>{
    if(stage.classList.contains('open')){
      if(ev.key==='Escape')closeThree();
      else if(ev.key==='ArrowDown'||ev.key==='ArrowRight'||ev.key==='PageDown'){cur=Math.min(cards.length-1,cur+1);place();ev.preventDefault()}
      else if(ev.key==='ArrowUp'||ev.key==='ArrowLeft'||ev.key==='PageUp'){cur=Math.max(0,cur-1);place();ev.preventDefault()}
      return;
    }
    if(ev.key==='Escape')panel.classList.remove('open')});
  // #t<token id> opens that word: the concordance links here.
  if(location.hash.startsWith('#t')){const id=decodeURIComponent(location.hash.slice(2));const w=document.querySelector('.w[data-id="'+CSS.escape(id)+'"]');if(w){w.scrollIntoView({block:'center'});show(id);concordLink(id)}}
})();
</script>
</body>
</html>`

const indexHTML = `<!doctype html>
<html lang="en">
<head>{{template "head"}}
<title>{{.Cfg.Title}}</title>
</head>
<body>
<header class="top"><div class="wrap"><a class="back" href="../index.html">← All works</a><h1>{{.Cfg.Title}} <span class="he" lang="he">{{.Cfg.He}}</span></h1></div></header>
<main class="wrap">
<p class="intro">A translation from the Hebrew with the analysis of every word: morphology, lemma, root, the lexicon's range of senses and the sense chosen, and the commentators anchored to the words they explain.</p>
{{range .Cfg.Sections}}<section class="seg"><a href="{{.Slug}}.html"><h1 style="font-size:22px">{{.Name}} <span class="he" lang="he">{{.He}}</span></h1></a><div class="ref">{{.From}}–{{.To}}</div></section>{{end}}
{{if .Missing}}<p class="intro">Not yet translated: {{range .Missing}}{{.}} {{end}}</p>{{end}}
</main>
<footer><div class="wrap"><ul>{{.Provenance}}</ul></div></footer>
<script>try{const t=localStorage.getItem('theme');if(t)document.documentElement.dataset.theme=t}catch(e){}</script>
</body>
</html>`
