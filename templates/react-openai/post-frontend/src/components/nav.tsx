import { Link } from 'react-router-dom';

const Nav = () => {
    return (
        <nav className="bg-slate-800 p-3">
            <Link to="/"
                className="text-white font-bold">
                Playground</Link>
        </nav>
    )
}

export default Nav;
