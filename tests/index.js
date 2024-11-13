import { sleep, check } from 'k6'
import http from 'k6/http'
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';

export const options = {
  vus: 10,
  duration: '30s'
}

/**
 * @method GET
 */
export async function get(url, params = undefined, headers = undefined) {
  const newParams = {
    headers: {
      'Content-Type': 'application/json',
      ...headers
    },
  }

  if (params) {
    const newUrl = new URL(url);
    Object.keys(params).forEach((i) => newUrl.searchParams.append(i, params[i]))
  
    return http.get(`${newUrl.toString()}`, newParams)
  }

  return http.get(`${url}`, newParams)
}

/**
 * @method POST
 */
export async function post(url, data) {
  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  }

  const payload = JSON.stringify(data)

  return http.post(`${url}`, payload, params)
}

const url = ''

async function doLogin() {
  const data = {
    email: '',
    password: ''
  }

  return await post(`${url}/auth/login`, data)
}

export default async function () {
  const resLogin = await doLogin()
  
  check(resLogin, { 'status of doLogin should be 201': (r) => r.status === 201 })

  sleep(1)
}
