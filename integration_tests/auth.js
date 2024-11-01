import http from 'k6/http';
import { check, fail } from 'k6';
import config from './config.js';
const auth_BASE_URL = `${config.base_url}/auth`;
const HEADERS = { headers: { 'Content-Type': 'application/json' } };

// export default function login() {
//     user_life_circle()
// }

// export function login(global_tag) {
//     const TAG = "user_life_circle"
// }

export function login(global_tag) {
    const tag = "login";
    const payload = JSON.stringify({ username: 'admin', password :'abc123' });
    const response = http.post(`${auth_BASE_URL}/login`, payload, HEADERS);
    const jsonResponse = response.json();

    const isValid = check(response, {
        [`${global_tag} ${tag} status is 200`]: (r) => r.status === 200,
        [`${global_tag} ${tag} response body is not empty`]: (r) => r.body.length > 0,
    }) && check(jsonResponse, {
        [`${global_tag} ${tag} success is true`]: (r) => r.success === true,
        [`${global_tag} ${tag} success is true`]: (r) => r.data.value,
    });

    if (!isValid) fail(`${tag} failed. Status: ${response.status}, Body: ${response.body}`);

    return jsonResponse.data.value;
}

