<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
{{template "inc/head.tpl" .}}
<div class="container">
  <div class="row">
{{if $.isLogin }}
<div class="col-lg-8 pull-right">
<form class="form-search">
<div class="col-xs-2">
<a href="/article/add" target="_blank" class="btn btn-success"> + New Blog</a>
</div>
<div class="col-xs-4">
 
 <select name="status" class="form-control">
      <option value="">选择状态</option>
      <option value="0">屏蔽</option>
      <option value="1">正常</option>
    </select>
	</div>
    <div class="input-group">
      <input type="text" name="title" class="form-control" placeholder="请输入标题">
      <span class="input-group-btn">
        <button class="btn btn-primary" type="submit">Search</button>
      </span>
    </div>
	</form>
  </div>
<div class="clearfix"></div>
{{end}}
	{{range $k,$v := .art}}
    <div class="media">
      <div class="media-body">
        <h4 class="media-heading"><a href="/article/{{$v.Id}}" title="{{$v.Title}}">{{$v.Title}}</a>（{{date $v.Created}}）</h4>
        <p>{{$v.Summary}}</p>
		{{if $.isLogin }}<p class="pull-right">{{if eq $v.Status 0}}<button type="button" class="btn btn-primary btn-xs">屏蔽</button>{{end}} <a href="/article/edit/{{$v.Id}}" class="btn btn-danger">修改</a></p>{{end}}
      </div>
    </div>
    {{end}}
    
    {{if .paginator.HasPages}}
    <ul class="pagination pagination">
      {{if .paginator.HasPrev}}
      <li><a href="{{.paginator.PageLinkFirst}}">首页</a></li>
      <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
      {{else}}
      <li class="disabled"><a>首页</a></li>
      <li class="disabled"><a>&laquo;</a></li>
      {{end}}
      {{range $index, $page := .paginator.Pages}} <li{{if $.paginator.IsActive .}} class="active"{{end}}> <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
      </li>
      {{end}}
      {{if .paginator.HasNext}}
      <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
      <li><a href="{{.paginator.PageLinkLast}}">尾页</a></li>
      {{else}}
      <li class="disabled"><a>&raquo;</a></li>
      <li class="disabled"><a>尾页</a></li>
      {{end}}
    </ul>
    {{end}}
  </div>
</div>
</body>
</html>
