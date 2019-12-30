import request from '../utils/request';

export const fetchData = query => {
    console.log('personal.js request')
    return request({
        url: './personal.json',
        method: 'get',
        params: query
    });
};

export function SavePersonal(formModel) {
    console.log('api.SavePersonal');
    const url = formModel.personId ? 'personal/edit' : 'personal/add'
    return request({
        url: url,
        method: 'post',
        data: formModel
    })
}
export function DeletePersonal(formModel) {
    console.log('DeletePersonal');
    return request({
        url: '/product/Save',
        method: 'post',
        data: formModel
    })
}
