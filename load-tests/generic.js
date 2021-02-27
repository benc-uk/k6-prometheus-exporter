//
// A generic load test script provided for example purposes and testing
//

import { check, sleep } from 'k6'
import http from 'k6/http'

const TARGET_URL = __ENV.TEST_TARGET || 'https://example.net'
const RAMP_TIME = __ENV.RAMP_TIME || '5m'
const RUN_TIME = __ENV.RUN_TIME || '10m'
const USER_COUNT = __ENV.USER_COUNT || 50
const SLEEP = __ENV.SLEEP || 0

// Very simple ramp up from zero to VUS_MAX over RAMP_TIME, then runs for further RUN_TIME
export let options = {
  stages: [
    { duration: RAMP_TIME, target: USER_COUNT },
    { duration: RUN_TIME, target: USER_COUNT },
  ],
}

// Totally generic HTTP check
export default function () {
  let res = http.get(TARGET_URL)

  check(res, {
    'Status is ok': (r) => r.status === 200,
  })

  sleep(SLEEP)
}
