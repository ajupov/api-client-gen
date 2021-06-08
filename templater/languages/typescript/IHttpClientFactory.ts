export interface IHttpClient {
    addHeaders: (headers?: HeadersInit) => void
    get: <Result>(url: string, parameters?: any) => Promise<Result>
    post: <Result>(url: string, parameters?: any, body?: any) => Promise<Result>
    put: <Result>(url: string, parameters?: any, body?: any) => Promise<Result>
    patch: <Result>(url: string, parameters?: any, body?: any) => Promise<Result>
    delete: <Result>(url: string, parameters?: any, body?: any) => Promise<Result>
}

export default interface IHttpClientFactory {
    readonly host: string
    createClient(host?: string): IHttpClient
}
