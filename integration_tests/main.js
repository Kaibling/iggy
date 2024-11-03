import user_checks from './users.js';
import workflow_checks from './workflows.js';


import { group } from 'k6';

export default function () {
    if (__ITER === 0) {
        // group('User endpoint - single request', () => {
        //     user_checks();
        //     if (__ITER == 0) {
        //         console.log("User endpoint check completed.");
        //     }
        // });
        group('Workflow endpoint - single request', () => {
            workflow_checks();
            if (__ITER == 0) {
                console.log("Workflow endpoint check completed.");
            }
        });
    }
}
