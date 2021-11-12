import { useEffect } from 'react';
import Calculator from './Calculator';
import initWasm from '../utils/initWasm';

export default function App() {
    useEffect(() => {
        initWasm()
    }, []);
    return (
        <div style={{ marginTop: 25 }}>
            <Calculator />
        </div>
    );
}
