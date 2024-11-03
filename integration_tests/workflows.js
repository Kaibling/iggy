import http from 'k6/http';
import { check, fail } from 'k6';
import config from './config.js';
import { login } from './auth.js';
const HEADERS = { headers: { 'Content-Type': 'application/json' } };
const WORKFLOW_BASE_URL = `${config.base_url}/workflows`;

export default function workflow_checks() {
    const token = login("workflow_checks");
    HEADERS.headers['Authorization'] = `Bearer ${token}`;
    //workflow_life_circle()
    workflow_shared_data_logs()
}

function workflow_life_circle() {
    const TAG = "workflow_life_circle"
    const id = add_workflow(TAG);
    add_existing_workflow(TAG);
    fetch_workflow(TAG, id);
    delete_workflow(id, TAG);
}

function add_workflow(global_tag) {
    const tag = "add_single_workflow";
    const payload = JSON.stringify({ name: 'test_workflow_1', code: "not going to be executed", object_type: "javascript" });
    const response = http.post(WORKFLOW_BASE_URL, payload, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);

    return jsonResponse.data.id;
}
function create_workflow(global_tag, payload) {
    const tag = "add_single_workflow";
    //const payload = JSON.stringify({ name: 'test_workflow_1',code:"not going to be executed", object_type: "javascript"  });
    const response = http.post(WORKFLOW_BASE_URL, payload, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);

    return jsonResponse.data.id;
}

function fetch_workflow(global_tag, id) {
    const tag = "fetch_workflow";
    const response = http.get(`${WORKFLOW_BASE_URL}/${id}`, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
        [`${global_tag} ${tag} check username`]: (r) => r.data.name == 'test_workflow_1',
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);
}

function add_existing_workflow(global_tag) {
    const tag = "add_existing_workflow";
    const payload = JSON.stringify({ name: 'test_workflow_1', code: "not going to be executed", object_type: "javascript" });
    const response = http.post(WORKFLOW_BASE_URL, payload, HEADERS);
    const jsonResponse = response.json();

    const isDuplicate = check(response, {
        [`${global_tag} ${tag} status is 500`]: (r) => r.status === 500,
        //[`${tag} status is 409`]: (r) => r.status === 409,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is false`]: (r) => r.success === false,
        [`${global_tag} ${tag} duplicate key error`]: (r) => r.data.includes("duplicate key"),
    });

    if (!isDuplicate) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);
}

function delete_workflow(id, global_tag) {
    const tag = "delete_workflow";
    const response = http.del(`${WORKFLOW_BASE_URL}/${id}`, null, HEADERS);

    if (!check(response, { [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200 })) {
        fail(`${global_tag} ${tag} failed. Status: ${response.status}, Body: ${response.body}`);
    }
}

function execute_workflow(id, global_tag) {
    const tag = "execute_workflow";
    const response = http.post(`${WORKFLOW_BASE_URL}/${id}/execute`, null, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);
}

function workflow_shared_data_logs() {
    const TAG = "workflow_shared_data_logs"

    // // create workflow that, saves data into shard_Data
    // const test_shared_workflow_1_payload = JSON.stringify({ name: 'test_shared_workflow_1', code: "shared_data['log_data'] = 'this data should be logged' ", object_type: "javascript" });
    // const test_shared_workflow_1_id = create_workflow(TAG, test_shared_workflow_1_payload);

    // // create workflow, that logs shared_data
    // const test_shared_workflow_2_payload = JSON.stringify({ name: 'test_shared_workflow_2', code: "log(shared_data['log_data'])", object_type: "javascript" });
    // const test_shared_workflow_2_id = create_workflow(TAG, test_shared_workflow_2_payload);

    // // create folder, thata execute both workflows above
    // const test_shared_folder_workflow_payload = JSON.stringify({ name: 'test_shared_folder_workflow', object_type: "folder" });
    // const test_shared_folder_workflow_id = create_workflow(TAG, test_shared_folder_workflow_payload);
    // execute folder
    const test_shared_folder_workflow_id = "01JBS5QJPVJS2RVQVKP4KQKX5Q"
    execute_workflow(test_shared_folder_workflow_id, TAG)


    // get run

    // get logs

    // check logs

    // delete everything

    // check run

    // check logs

}
