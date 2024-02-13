const Recipe = () => {
    return <div className="w-3/4 mx-auto">
        <div className="recipe-btns flex justify-between">
            <div className="recipe-btns-left flex">
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 mr-4">Main</div>
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600">Dessert</div>
            </div>
            <div className="recipe-btns-right flex">
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 mr-4">Search</div>
                <div className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600">+ Add new recipe</div>
            </div>
        </div>
        <div>recipes</div>
    </div>
    ;
};

export default Recipe;