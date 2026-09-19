// Q: Write usePrevious(value) that returns the value from the previous render.
import { useRef, useEffect } from "react";

export function usePrevious(value) {
    const ref = useRef();
    useEffect(() => {
        ref.current = value; // runs after render, so the return below is still the old value
    });
    return ref.current;
}
