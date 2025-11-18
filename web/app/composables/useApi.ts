type useFetchType = typeof useFetch

export const useApi: useFetchType = (path, options = {}) => {
    const config = useRuntimeConfig()

    options.baseURL = config.public.serverBaseURL;

    return useFetch(path, options)
}
