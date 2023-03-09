export type ApiResponse = {
    error: boolean
    code: number
    message: unknown
}

export const apiBase = 'http://localhost:8000'