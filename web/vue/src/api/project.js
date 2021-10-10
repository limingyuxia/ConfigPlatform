import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/project/list',
    method: 'get',
    params
  })
}

export function upList(params, type) {
  console.log('upList:', params, type)
  /*
  return request({
    url: '/project/list',
    method: 'get',
    params
  })*/
}

export function getDetailList(params, type) {
  console.log('getDetailList:', params, type)

  return request({
    url: '/project/detail',
    method: 'get',
    params
  })
}
export function addProject(params, type) {
  console.log('addProject:', params, type)

  return request({
    url: '/project/add',
    method: 'post',
    data: params
  })
}

export function deleteProject(params, type) {
  console.log('deleteProject:', params, type)

  return request({
    url: '/project/delete',
    method: 'DELETE',
    params
  })
}

export function editProject(params, type) {
  console.log('editProject:', params, type)

  return request({
    url: '/project/edit',
    method: 'POST',
    data: params
  })
}

