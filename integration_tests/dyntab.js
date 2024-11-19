import http from 'k6/http';
import { check, fail } from 'k6';
import config from './config.js';
const DYN_TAB_BASE_URL = `${config.base_url}/dynamic-tables`;
const HEADERS = { headers: { 'Content-Type': 'application/json' } };
const global_tag = "dyn_tab"

export default function check_dyn_tabs() {
    dyntab_life_circle()
}

function dyntab_life_circle() {
    const TAG = "user_life_circle"

        // create table
    const id = create_table(TAG);

    // add variable
    // show table
    // remove variable
    // delete table
}

function create_table(global_tag) {
    const tag = "create_dyn_Table";
    const payload = JSON.stringify([{ name: 'test_table'}]);
    const response = http.post(`${DYN_TAB_BASE_URL}`, payload, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
        [`${global_tag} ${tag} data exists`]: (r) => r.data,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);

    return jsonResponse.data[0].id;
}

function add_variables(global_tag,dyn_tag_id) {
    const tag = "add_dyn_Table";

    const payload = JSON.stringify([{ name: 'id', variable_type: 'dinge',dynamic_table_id : dyn_tag_id}]);
    const response = http.post(`${DYN_TAB_BASE_URL}`, payload, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
        [`${global_tag} ${tag} data exists`]: (r) => r.data,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);

    return jsonResponse.data[0].id;
}


