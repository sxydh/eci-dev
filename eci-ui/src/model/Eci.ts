import {Container} from '@/model/Container';

export interface Eci {
    id: string;
    name: string;
    replicaName: string;
    ready: string;
    status: string;
    restarts: string;
    age: string;
    ip: string;
    node: string;
    containers: Container[];
}