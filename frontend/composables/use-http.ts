import { FetchOptions, $fetch } from 'ofetch'
import { useAuthStore } from '../store/auth'


const opts: FetchOptions = {
    async onRequest({ request, options }) {
        const auth = useAuthStore()
        options.baseURL = 'http://localhost:8000/'

        if (options.body instanceof Object) {
            options.body = JSON.stringify(options.body)
        }
        if (auth.token != "") {
            options.headers = options.headers || {}
            options.headers['Authorization'] = 'Bearer ' + auth.token
        }
    }
}

export default () => $fetch.create(opts)