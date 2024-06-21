import {Resource} from '@/model/Resource';

export interface Container {
    name: string;
    image: string;
    imagePullPolicy: string;
    command: string;
    ready: string;
    restarts: string;
    resourceRequest: Resource;
    resourceLimit: Resource;
}