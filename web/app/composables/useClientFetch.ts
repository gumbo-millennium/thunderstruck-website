export const useClientFetch = (path: string, options = {}) => {
    const config = useRuntimeConfig()
    return $fetch(`${config.public.clientBaseURL}${path}`, options)
}
