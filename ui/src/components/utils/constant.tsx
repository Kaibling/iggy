import { Field } from './types.tsx'

export const WorkflowFields: Field[] = [
    { fieldName: "name", displayName: "Name", collection: "workflows" },
    { fieldName: "object_type", displayName: "Object type",symbol: true  },
    { fieldName: "fail_on_error", displayName: "Fail on error", boolean: true },
    { fieldName: "created_at", displayName: "Created At",date:true },
    { fieldName: "created_by", displayName: "Created By" },
    { fieldName: "modified_at", displayName: "Modified At",date:true },
    { fieldName: "modified_by", displayName: "Modified By" }
];


export const WorkflowFieldsSmall: Field[] = [
    { fieldName: "name", displayName: "Name", collection: "workflows" },
    { fieldName: "object_type", displayName: "Object type", symbol: true },
    { fieldName: "fail_on_error", displayName: "Fail on error", boolean: true }
];


export const RunFields: Field[] = [
    { fieldName: "id", displayName: "id", collection: "runs" },
    { fieldName: "workflow", displayName: "workflow", collection: "workflows", identifier : "workflow" },
    { fieldName: "error", displayName: "Error" },
    { fieldName: "result", displayName: "Result" },
    { fieldName: "run_time", displayName: "run_time" },
    { fieldName: "start_time", displayName: "Starttime" ,date:true },
    { fieldName: "finish_time", displayName: "Endtime",date:true },
    { fieldName: "created_by", displayName: "Run by" },
];