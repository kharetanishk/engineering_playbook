// Q: Write useFetch(url) returning { data, loading, error }. Handle unmount / url change races.
import { useState, useEffect } from "react";

export function useFetch(url) {
    const [state, setState] = useState({ data: null, loading: true, error: null });

    useEffect(() => {
        const ctrl = new AbortController();
        setState({ data: null, loading: true, error: null });

        fetch(url, { signal: ctrl.signal })
            .then((r) => {
                if (!r.ok) throw new Error(`HTTP ${r.status}`);
                return r.json();
            })
            .then((data) => setState({ data, loading: false, error: null }))
            .catch((error) => {
                if (error.name !== "AbortError") setState({ data: null, loading: false, error });
            });

        return () => ctrl.abort(); // cancels stale request
    }, [url]);

    return state;
}
