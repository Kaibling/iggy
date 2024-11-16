export interface Workflow {
    id: string,
    name: string;
    code: string;
    object_type: string;
    fail_on_error: boolean;
    children: Workflow[];
    meta: Meta;

}

export interface Meta {
    created_at: string;
    created_by: string;
    modified_at: string;
    modified_by: string;
}


export interface NewWorkflow {
    name: string;
    code: string;
    object_type: string;
    fail_on_error: boolean;
}


export interface Field {
    fieldName: string;
    displayName: string;
    boolean?: boolean;
    date?: boolean;
    collection?: string;
    identifier?: string;
    symbol?: boolean;
}

export interface Option {
    id: string;
    name: string;
}
export interface Identifier {
    id: string;
    name: string;
}



export interface Run {
    id: string,
    workflow: Identifier;
    error: string;
    result: string;
    run_time: string;
    start_time: string;
    finish_time: string;
    meta: Meta;
}


// types.ts
export interface Workflow2 {
    id: string;
    name: string;
    code: string;
    object_type: string;
    children: Workflow2[] | null;
  }
  