$(function () {
    search.init(true);
})

var search = {
    init: function (isInitGrid) {
        $.ajax({
            url: "/config/getSearchOptions",
            type: 'GET',
            dataType: "json",
            success: function (data) {
                if (data.Profiles) {
                    var profile = $("#searchForm select[name='profile']");
                    var _html = [];
                    _html.push('<option value="">请选择</option>')
                    for (var i = 0; i < data.Profiles.length; i++) {
                        _html.push('<option value="' + data.Profiles[i] + '">' + data.Profiles[i] + '</option>');
                    }
                    profile.html(_html.join(""));
                    $("#profile").html(_html.join(""));
                    $("#currentProfile").html(_html.join(""));
                }
                if (data.Applications) {
                    var application = $("#searchForm select[name='application']");
                    var _html = [];
                    _html.push('<option value="">请选择</option>')
                    for (var i = 0; i < data.Applications.length; i++) {
                        _html.push('<option value="' + data.Applications[i] + '">' + data.Applications[i] + '</option>');
                    }
                    application.html(_html.join(""));
                    $("#application").html(_html.join(""));
                }
                if (isInitGrid) {
                    grid.init();
                }
            }
        });
    }
}

var grid = {
    init: function () {
        var columns = (function () {
            var cs = [];
            var checkboxField = {
                checkbox: true
            }
            cs.push(checkboxField);
            var configIdField = {
                field: 'ConfigId',
                title: 'ID',
                align: 'center'
            }
            cs.push(configIdField);
            var configKeyField = {
                field: 'ConfigKey',
                title: '配置键',
            }
            cs.push(configKeyField);
            var configValueField = {
                field: 'ConfigValue',
                title: '配置值',
                align: 'center'
            }
            cs.push(configValueField);
            var applicationField = {
                field: 'Application',
                title: '项目服务名称',
                align: 'center'
            }
            cs.push(applicationField);
            var profileField = {
                field: 'Profile',
                title: '环境',
                align: 'center'
            }
            cs.push(profileField);
            var labelField = {
                field: 'Label',
                title: '版本',
                align: 'center'
            }
            cs.push(labelField);
            var configDescField = {
                field: 'ConfigDesc',
                title: '描述',
                align: 'center'
            }
            cs.push(configDescField);
            return cs;
        })();

        $("#table").bootstrapTable({
            url: "/config/list",
            height: 580,
            striped: true,
            columns: columns,
            // detailView: true,//父子表
            pagination: true,
            sidePagination: 'server',
            clickToSelect: true,
            toolbar: "#toolbar",
            queryParamsType: '',
            queryParams: function (params) {
                $.each($("#searchForm").serializeArray(), function (i, field) {
                    params[field.name] = field.value;
                });
                return params;
            },
            responseHandler: function (res) {
                return {total: res.total, rows: res.rows}
            }
        });
        // 数据加载成功后 进行toolbar按钮事件绑定
        grid.toolbarEvent();
        // 搜索栏构造及查询按钮事件绑定
        grid.searchForm();
    },
    searchForm: function () {
        var searchBtn = $("#search");
        searchBtn.on("click", function () {
            var data = {};
            $.each($("#searchForm").serializeArray(), function (i, field) {
                data[field.name] = field.value;
            });

            data.pageNumber = 1;
            var params = {
                silent: true,
                query: data
            }
            $('#table').bootstrapTable("refresh", params);
            return false;
        });

    },
    toolbarEvent: function () {
        // 新增
        $("#add").on("click", function () {
            $('#myModalLabel').html("新增配置");
            $(".batchEditForm").hide();
            $(".editForm").show();
            $('#myModal').modal();
            $("#editForm form")[0].reset();
            $("#configId").val("");
        });
        $("#save").on("click", function () {
            var data = {};
            $.each($("#editForm").find("form").serializeArray(), function (i, field) {
                data[field.name] = field.value;
            });
            var url = "/config/add";
            if (data.configId && url.configId != "") {
                url = "/config/update";
            }
            $.ajax({
                url: "" + url,
                // async : false,
                type: 'POST',
                dataType: "json",
                data: data,
                success: function (data) {
                    console.log(data)
                    if (data) {
                        if (data.flag == 1) {
                            alert('操作成功!');
                            $('#myModal').modal('hide');
                            $('#table').bootstrapTable("refresh");
                        } else {
                            alert(data.msg);
                        }
                    } else {
                        alert('操作失败!');
                    }
                }
            });
            return false;
        });
        // 修改
        $("#update").on("click", function () {
            var rows = $('#table').bootstrapTable("getSelections");
            if (rows.length != 1) {
                alert('选择一条记录!');
                return;
            }
            $('#myModalLabel').html("修改配置");
            $(".batchEditForm").hide();
            $(".editForm").show();
            $('#myModal').modal();
            $("#editForm form")[0].reset();
            $("#configId").val(rows[0].ConfigId);
            $("#profile").val(rows[0].Profile);
            $("#application").val(rows[0].Application);
            $("#configKey").val(rows[0].ConfigKey);
            $("#configValue").val(rows[0].ConfigValue);
            $("#configDesc").val(rows[0].ConfigDesc);
        });
        // 删除
        $("#remove").on("click", function () {
            var rows = $('#table').bootstrapTable("getSelections");
            if (rows.length > 0) {
                var keys = [];
                var ids = [];
                for (var i = 0; i < rows.length; i++) {
                    ids.push(rows[i].ConfigId);
                    keys.push(rows[i].ConfigKey)
                }
                if (confirm("确认删除配置【" + keys.join(",") + "】")) {
                    // 删除
                    $.ajax({
                        url: "/config/delete",
                        // async : false,
                        type: 'POST',
                        dataType: "json",
                        contentType: 'application/json',
                        data: JSON.stringify(ids),
                        success: function (data) {
                            if (data && data.flag) {
                                alert('删除成功!');
                                $('#table').bootstrapTable("refresh");
                            } else {
                                alert('删除失败!');
                            }
                        }
                    });
                }
            } else {
                alert('至少选择一条记录!');
            }
            return false;
        });
        // 修改环境
        $("#updateProfile").on("click", function () {
            $('#myModalLabel').html("修改环境");
            $(".editForm").hide();
            $(".batchEditForm").show();
            $('#myModal').modal();
            $("#batchEditForm form")[0].reset();
        });
        $("#batchSave").on("click", function () {
            var currentProfile = $("#currentProfile").val();
            var deployProfile = $("#deployProfile").val();
            if (!currentProfile) {
                alert("当前环境不能为空！");
                return ;
            }
            if (!deployProfile) {
                alert("部署环境不能为空！");
                return ;
            }
            $.ajax({
                url: "/config/batchUpdateProfile/" + currentProfile + "/" + deployProfile,
                // async : false,
                type: 'POST',
                dataType: "json",
                success: function (data) {
                    if (data) {
                        if (data.flag == 1) {
                            alert('操作成功!');
                            $('#myModal').modal('hide');
                            search.init(false);
                            $('#table').bootstrapTable("refresh");
                        } else {
                            alert(data.msg);
                        }
                    } else {
                        alert('操作失败!');
                    }
                }
            });
            return false;
        });
    }
}