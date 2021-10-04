const Mock = require('mockjs')

const data = Mock.mock({
  'items|30': [{
    id: '@id',
    title: '@sentence(10, 20)',
    'status|1': ['published', 'draft', 'deleted'],
    author: 'name',
    display_time: '@datetime',
    pageviews: '@integer(300, 5000)'
  }]
})

module.exports = [
  /*{
    url: '/config/list',
    type: 'get',
    response: config => {
      const items = data.items
      return {
        "list": [
          {
            "admin": "1",
            "c_time": "2",
            "create_time": "3",
            "department": "4",
            "description": "5",
            "id": 0,
            "name": "6",
            "u_time": "7",
            "update_time": "8"
          }
        ],
        "total": 1
      }
    }
  }*/
]
