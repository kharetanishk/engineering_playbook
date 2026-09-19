// Q: Write useDebounce(value, delay) that returns the value only after it stops changing for `delay` ms.
import { useState, useEffect } from "react";

export function useDebounce(value, delay = 300) {
    const [debounced, setDebounced] = useState(value);

    useEffect(() => {
        const id = setTimeout(() => setDebounced(value), delay);
        return () => clearTimeout(id); // reset timer on every change / unmount
    }, [value, delay]);

    return debounced;
}
