import request from '@/util/request';
import {Eci} from '@/model/Eci';

export function getEciList() {
    return request.get<Eci[]>('eci/list');
}

export function addEci(body: Object) {
    return request.post('eci', body);
}

export function deleteEci(name: string) {
    return request.delete('eci', {
        params: {
            name: name
        }
    });
}