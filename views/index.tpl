<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8"/>
    <title>配置管理</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/static/css/bootstrap-table.min.css"/>
    <style type="text/css">
        body {
            background: #F8F8F8;
        }
        .main {
            width: 96%;
            margin-left: auto;
            margin-right: auto;
            padding-top: 10px;
        }
    </style>
</head>
<body>
<!--表格-->
<div class="main">
    <div id="toolbar">
        <button id="add" class="btn btn-primary">新增</button>
        <button id="update" class="btn btn-warning">修改</button>
        <button id="remove" class="btn btn-danger">删除</button>
        <button id="updateProfile" class="btn btn-danger">修改环境</button>
    </div>
    <form id="searchForm" class="form-inline" role="form" >
        <div class="form-group" style="margin-right: 15px;">
            <label class="form-label">环境:</label>
            <select name="profile" class="form-control"></select>
        </div>
        <div class="form-group" style="margin-right: 15px;">
            <label class="form-label">项目服务名称:</label>
            <select name="application" class="form-control"></select>
        </div>
        <div class="form-group" style="margin-right: 15px;">
            <label class="form-label">配置键:</label>
            <input type="text" class="form-control" name="configKey"/>
        </div>
        <div class="form-group" style="margin-right: 15px;">
            <label class="form-label">配置值:</label>
            <input type="text" class="form-control" name="configValue"/>
        </div>
        <button id="search" class="btn btn-default">查询</button>
    </form>
    <table id="table"></table>
</div>

<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <h4 class="modal-title" id="myModalLabel">新增</h4>
            </div>
            <div class="modal-body">
                <!--表单-->
                <div id="editForm" class="editForm">
                    <form class="form-horizontal" role="form">
                        <input name="configId" id="configId" type="hidden"/>
                        <div class="form-group">
                            <label for="profile" class="col-sm-2 control-label">环境:</label>
                            <div class="col-sm-10">
                                <select class="form-control" name="profile" id="profile"/>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="application" class="col-sm-2 control-label">项目服务名称:</label>
                            <div class="col-sm-10">
                                <select class="form-control" name="application" id="application"/>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="configKey" class="col-sm-2 control-label">配置键:</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" name="configKey" id="configKey"/>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="configValue" class="col-sm-2 control-label">配置值:</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" name="configValue" id="configValue"/>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="configDesc" class="col-sm-2 control-label">描述:</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" name="configDesc" id="configDesc"/>
                            </div>
                        </div>
                    </form>
                </div>

                <!--批量修改环境-->
                <div id="batchEditForm" class="batchEditForm">
                    <form class="form-horizontal" role="form">
                        <div class="form-group">
                            <label for="currentProfile" class="col-sm-2 control-label">当前环境:</label>
                            <div class="col-sm-10">
                                <select class="form-control" name="currentProfile" id="currentProfile"/>
                            </div>
                        </div>
                        <div class="form-group">
                            <label for="deployProfile" class="col-sm-2 control-label">部署环境:</label>
                            <div class="col-sm-10">
                                <input type="text" class="form-control" name="deployProfile" id="deployProfile"/>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                <button type="button" id="save" class="btn btn-primary editForm">提交</button>
                <button type="button" id="batchSave" class="btn btn-primary batchEditForm">提交</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>

<script type="application/javascript" src="/static/js/jquery-1.9.1.min.js"></script>
<script type="application/javascript" src="/static/js/bootstrap.min.js"></script>
<script type="application/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="application/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
<script type="application/javascript" src="/static/js/json2-min.js"></script>
<script type="application/javascript" src="/static/js/index.js"></script>
</body>
</html>