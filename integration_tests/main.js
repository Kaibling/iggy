import user_checks from './users.js';
// import {fetch_all_user} from './users.js'
import { group } from 'k6';
//Configure load stages, thresholds, and custom settings
// export let options = {
//     stages: [
//         { duration: '30s', target: 10 }, // Ramp-up to 10 VUs in 30s
//         { duration: '1m', target: 30 },  // Sustain 30 VUs for 1 minute
//         { duration: '30s', target: 50 }, // Further ramp-up to 50 VUs
//         { duration: '1m', target: 50 },  // Hold at 50 VUs for 1 minute
//         { duration: '10s', target: 0 },  // Ramp-down
//     ],
//     thresholds: {
//         // 95% of requests should complete under 400ms
//         http_req_duration: ['p(95)<400'],
//         // Less than 2% of requests should fail
//         http_req_failed: ['rate<0.02'],
//         // Ensure system error rate stays below 1%
//         checks: ['rate>0.99'],
//     },
//     rps: 30, // 30 requests per second limit
// };

export default function () {
    if (__ITER === 0) {
        group('User endpoint - single request', () => {
            user_checks();
            if (__ITER == 0) {
                console.log("User endpoint check completed.");
            }
        });
        // group('Workflow endpoint - single request', () => {
        //     workflow_checks();
        //     if (__ITER == 0) {
        //         console.log("Workflow endpoint check completed.");
        //     }
        // });
    }
    // group('user endpoint - loadtest request', () => {
    //     fetch_all_user("fetch_all_user");
    // });
}
