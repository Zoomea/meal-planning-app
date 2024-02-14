import siteLogo from './assets/siteLogo.png';
import downArrow from './assets/down-arrow.svg';

const Recipe = () => {
    return <div className="w-full h-48 bg-green-300 rounded-xl hover:shadow-lg transition">
        <img className="w-full h-4/5 rounded-t-xl" src={siteLogo} alt="recipe pic" />
        <div className="mx-4 h-1/5 flex justify-between items-center text-green-700">
            <div>Name</div>
            <img className="w-2" src={downArrow} alt="down arrow" />
        </div>
    </div>
}

export default Recipe;

