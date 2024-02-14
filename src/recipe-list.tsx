import Recipe from './recipe.tsx';

const Recipes = () => {
    return <div className="w-3/4 mx-auto">
        <div className="recipe-btns flex justify-between">
            <div className="recipe-btns-left flex">
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 mr-4 hover:shadow-lg transition">Main</div>
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 hover:shadow-lg transition">Dessert</div>
            </div>
            <div className="recipe-btns-right flex">
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 mr-4 hover:shadow-lg transition">Search</div>
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 hover:shadow-lg transition">+ Add new recipe</div>
            </div>
        </div>
        {/* get the data from backend and use map to show each one */}
        <div className="w-full grid grid-cols-1 sm:grid-cols-3 md:grid-cols-4 gap-8 my-4">
            <Recipe />
            <Recipe />
            <Recipe />
            <Recipe />
            <Recipe />
            <Recipe />
            <Recipe />
            <Recipe />
        </div>
        
    </div>
    ;
};

export default Recipes;