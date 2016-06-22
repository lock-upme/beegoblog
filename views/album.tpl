<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>相册-{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/js/lightbox/css/lightbox.min.css" media="all" rel="stylesheet" type="text/css" />
</head>
<body>
{{template "inc/head.tpl" .}}
<div class="container">
  <div class="row">
{{if $.isLogin }}
<div class="col-lg-8 pull-right" style="margin-bottom:10px;">
<form class="form-search">
<div class="col-xs-2">
<a href="/album/upload" target="_blank" class="btn btn-success"> + 上传图片</a>
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
	{{range $k,$v := .alb}}
    
	<div class="col-sm-6 col-md-4">
        <div class="thumbnail">
          <a href="{{$v.Picture}}" data-lightbox="example-set" data-title="{{$v.Summary}}"><img alt="{{$v.Title}}" style="height: 200px; width: 100%; display: block;" src="{{$v.Picture}}"></a>
          <div class="caption">
            <h3>{{$v.Title}}</h3>
            <p>{{substr $v.Summary 0 20}}&nbsp;</p>
			{{if $.isLogin }}
            <p><a href="javascript:;" class="btn btn-primary js-album-edit" data-id="{{$v.Id}}" data-title="{{$v.Title}}" data-summary="{{$v.Summary}}" data-status="{{$v.Status}}">修改</a> <a href="javascript:;" class="btn btn-default">{{if $v.Status}}正常{{else}}屏蔽{{end}}</a></p>{{end}}
          </div>
        </div>
      </div>	
    {{end}}
    
    {{if .paginator.HasPages}}
	<div class="clearfix"></div>
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
{{template "inc/foot.tpl" .}}
<script src="/static/js/lightbox/js/lightbox.min.js" type="text/javascript"></script>
</body>
</html>
