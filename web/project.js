function AddProject() {
    window.alert("添加成功")
}

function GetProjectDetail() {
    window.alert("编辑成功");
}

function DeleteProject() {
    window.confirm('确认删除吗?')
}

function displayProjectInfo(projectInfo) {
    var tableElem = window.document.getElementById("InfoTable")
    
    var data = ""
    var userInfos = projectInfo.list

    // 编辑调用detail，打开新界面，画一个表单
    // 删除是确认弹窗

    for (var i = 0; i < userInfos.length; i++) {
        data += 
            "<tr align=\"center\">"                     +
            "<td>" + userInfos[i].name + "</td>"        +
            "<td>" + userInfos[i].description + "</td>" +
            "<td>" + userInfos[i].department + "</td>"  +
            "<td>" + userInfos[i].admin + "</td>"       +
            "<td>" + userInfos[i].c_time + "</td>"      +
            "<td>" + userInfos[i].u_time + "</td>"      +
            "<td>" + "<input type=\"button\" value=\"编辑\" onclick=\"GetProjectDetail()\">"
                   + "<input type=\"button\" value=\"删除\" onclick=\"DeleteProject()\">" + "</td>" +
            "</tr>"

        window.document.getElementById("content").innerHTML = data
    }   
    
    window.document.getElementById("total").innerText = 
        "总数:" + projectInfo.total
}


function GetProjectList() {
    var httpRequest = new XMLHttpRequest()
    var url = "/config/list"

    httpRequest.open('GET', url, true)
    httpRequest.setRequestHeader("Content-type", "application/json")

    var body = {}
    httpRequest.send(JSON.stringify(body))

    httpRequest.onreadystatechange = function () {
        //验证请求是否发送成功
        if (httpRequest.readyState == 4 && httpRequest.status == 200) {
            var json = httpRequest.responseText;
            var obj = JSON.parse(json)
            
            displayProjectInfo(obj);
        }
    }
}