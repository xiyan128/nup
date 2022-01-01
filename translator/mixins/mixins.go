package mixins

const Caption = `
{{if or .numbered .caption }}
	<div class="caption-wrapper">
    {{if .numbered}}<a class="caption-num" href="#{{.id}}">{{._caption_type}} {{._caption_num}}</a>{{end}}
	{{if .caption}}<div class="caption">{{.caption}}</div>{{end}}
	</div>
{{end}}`

const MathNumber = `
	{{if .numbered}}
	    <a class="math-num">{{._caption_num}}</a>
	{{end}}
`
