export interface IHttpClient {
    get: <Result>(url: string, data?: any, headers?: HeadersInit) => Promise<Result>
    post: <Result>(url: string, data?: any, headers?: HeadersInit) => Promise<Result>
    put: <Result>(url: string, data?: any, headers?: HeadersInit) => Promise<Result>
    patch: <Result>(url: string, data?: any, headers?: HeadersInit) => Promise<Result>
    delete: <Result>(url: string, data?: any, headers?: HeadersInit) => Promise<Result>
}

export default interface IHttpClientFactory {
    readonly host: string
    createClient(host?: string): IHttpClient
}
