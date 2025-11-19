export const useClientFetch = <T>(path: string, options = {}) => {
    const config = useRuntimeConfig()
    return $fetch<T>(`${config.public.clientBaseURL}${path}`, options)
}
