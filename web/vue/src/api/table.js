import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/config/list',
    method: 'get',
    params
  })
}

export function getConfingList(params) {
  console.log('获取数据表格:', params)
  console.log('数据:', params.dataAll)
  console.log('头:', params.header)

  return request({
    url: '/table/getOpticalList.php',
    method: 'post',
    data: params.dataAll,
    headers: params.header

  })
}
export async function checkOpticalList(params) {
  console.log('上传:', params)

  params.header['Content-Type'] = 'application/x-www-form-urlencoded'

  return request({
    url: '/table/checkOpticalList.php',
    method: 'post',
    data: params.dataAll,
    headers: params.header

  })
}

export async function upOpticalList(params) {
  console.log('上传:', params)

  params.header['Content-Type'] = 'application/x-www-form-urlencoded'

  return request({
    url: '/table/upOpticalList.php',
    method: 'post',
    data: params.dataAll,
    headers: params.header

  })
}
