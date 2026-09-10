interface apiResponse {
    ok: boolean
    response: object | null
}

interface apiService {
    apiCall: (path: string, method: string, payload?: object) => Promise<apiResponse>
    loading: boolean
}

export function newApiService(): apiService {
    const service: apiService = {
        loading: true,
        apiCall: async(path: string, method: string, payload?: object): Promise<apiResponse> => {
            try {
                const respose = await fetch(path, {
                    method: method,
                    credentials: "include",
                    body: JSON.stringify(payload)
                })

                if (!respose.ok) {
                    return {ok: false, response: null}
                }

                const result = await respose.json()

                service.loading = false

                return {ok: true, response: result}
            } catch(ex) {
                service.loading = false
                return {ok: false, response: null}
            }
        }
    }

    return service
}