import { Link } from "react-router-dom";
import TopLogo from "../../talwar-logos_transparent.png";

const TopBar = () => {
    return (
        <>
            <nav className="bg-white sticky top-0 z-20 shadow-sm">
                <div className="container mx-auto flex justify-between items-center  ">

                    <div className="flex items-center">
                        <Link to="/">
                            <div className="w-[80px] relative">
                                <img alt="Talwar" src={TopLogo} />
                            </div>
                        </Link>
                        <div className="ml-14 flex items-center gap-9 text-secondary-400 font-semibold">

                        </div>
                    </div>
                    <div className="flex flex-row space-x-10 font-bold  italic text-lg">
                        {/* <Link to="/history">History 📜</Link>
                        <a href="/">Quizer 🙋</a>
                        <a href="/leaderboard">LeaderBoard 🏆</a> */}
                    </div>
                    <div className="flex items-center gap-5">


                        <Link to="/add" className="flex py-3 px-5 text-secondary-400">
                            <span className="font-bold mr-2">Add target</span>
                        </Link>
                        <Link to="/schedule" className="flex py-3 px-5 text-secondary-400">
                            <span className="font-bold mr-2">Scan Schedules</span>
                        </Link>
                        <Link to="/email-settings" className="flex py-3 px-5 text-secondary-400">
                            <span className="font-bold mr-2" >Email Settings</span>
                        </Link>


                    </div>
                </div>
            </nav>
        </>
    )
}

export default TopBar;