import request from '../utils/request';

export const fetchData = query => {

    return request({
        url: './adjuest.json',
        method: 'get',
        params: query
    });
};
