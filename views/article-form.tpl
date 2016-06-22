<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>博客编辑-{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
</head>
<body>
{{template "inc/head.tpl" .}}
<div class="container">
  <div class="row">
    <form method="post" id="article-form">
      <div class="form-group">
        <label for="title">Title{{.ss}}</label>
        <input type="text" class="form-control" id="title" name="title" placeholder="blog title" value="{{.art.Title}}">
      </div>
      <div class="form-group">
        <label for="uri">URL</label>
        <input type="url" class="form-control" id="uri" name="uri" placeholder="http://" value="{{.art.Uri}}">
      </div>
      <div class="form-group">
        <label for="author">Author</label>
        <input type="text" class="form-control" id="author" name="author" placeholder="Author" value="{{.art.Author}}">
      </div>
      <div class="form-group">
        <label for="keywords">keywords</label>
        <input type="text" class="form-control" id="keywords" name="keywords" placeholder="Keywords" value="{{.art.Keywords}}">
      </div>
      <div class="form-group">
        <label for="summary">Summary</label>
        <textarea class="form-control" id="summary" name="summary" placeholder="Summary" style="height:60px;">{{.art.Summary}}</textarea>
      </div>
      <div class="form-group">
        <label for="content">Content</label>
        <textarea class="form-control" id="blogContent" name="content" placeholder="Content" style="height:300px;">{{.art.Content}}</textarea>
      </div>
      <div class="form-group">
        <label for="content">Status</label>
        <label class="radio-inline">
        <input type="radio" name="status" value="1" {{if eq .art.Status 1}}checked{{end}}>
        正常 </label>
        <label class="radio-inline">
        <input type="radio" name="status" value="0" {{if eq .art.Status 0}}checked{{end}}>
        屏蔽 </label>
      </div>
      <input type="hidden" name="id" value="{{.art.Id}}" />
      <button type="button" onClick="history.back();" class="btn btn-default">Back</button>
      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</div>
{{template "inc/foot.tpl" .}}
<script src="/static/keditor/kindeditor-min.js"></script>
<script>
$(function(){
	var editor = KindEditor.create("#blogContent", {
	    uploadJson: "/upload",
	    fileManagerJson: "/board/keditor/filemanager",
	    allowFileManager: true,
	    filterMode : false,
	    afterBlur: function(){this.sync();}
	});	
})
</script>
</body>
</html>
