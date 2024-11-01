import http from 'k6/http';
import { check, fail } from 'k6';
import config from './config.js';
import {login} from './auth.js';
const HEADERS = { headers: { 'Content-Type': 'application/json' } };
const USER_BASE_URL = `${config.base_url}/users`;
export default function user_checks() {
    user_life_circle()
}

function user_life_circle() {
    const TAG = "user_life_circle"
    const token = login(TAG);
    HEADERS.headers['Authorization'] = `Bearer ${token}`;
    const id = add_user(TAG);
    add_existing_user(TAG);
    fetch_user(TAG,id);
    delete_user(id,TAG);
}

function add_user(global_tag) {
    const tag = "add_single_user";
    const payload = JSON.stringify({ username: 'test_1' });
    const response = http.post(USER_BASE_URL, payload, HEADERS);
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
function fetch_user(global_tag,id) {
    const tag = "fetch_user";
    const response = http.get(`${USER_BASE_URL}/${id}`, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
        [`${global_tag} ${tag} check username`]: (r) => r.data.username == 'test_1',
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);
}
export function fetch_all_user(global_tag) {
    const tag = "fetch_all_user";
    const response = http.get(USER_BASE_URL,null, HEADERS);
    //const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);
}

function add_existing_user(global_tag) {
    const tag = "add_existing_user";
    const payload = JSON.stringify({ username: 'test_1' });
    const response = http.post(USER_BASE_URL, payload, HEADERS);
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

function delete_user(id,global_tag) {
    const tag = "delete_user";
    const response = http.del(`${USER_BASE_URL}/${id}`,null,HEADERS);

    if (!check(response, { [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200 })) {
        fail(`${global_tag} ${tag} failed. Status: ${response.status}, Body: ${response.body}`);
    }
}
