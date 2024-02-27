import { useState, ChangeEvent } from "react";
import StarImg from './assets/star.svg';

const NewRecipeModal = () => {

    const [titleValue, setTitleValue] = useState<string>('');
    const [descriptionValue, setDescriptionValue] = useState<string>('');
    const [ingredientsValue, setIngredientsValue] = useState<string>('');

    const handleTitleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setTitleValue(e.target.value);
    }

    const handleDescriptionChange = (e: ChangeEvent<HTMLInputElement>) => {
        setDescriptionValue(e.target.value);
    }

    const handleIngredientsChange = (e: ChangeEvent<HTMLInputElement>) => {
        setIngredientsValue(e.target.value);
    }

    return <div className="recipe-modal-container">
        <div className="text-green-500 text-xl mb-4">New recipe</div>
        <div className="flex">
            <div className="w-3/5">
                <div className="mb-2">
                    <label htmlFor="title">Title:</label>
                    <input
                        type="text"
                        id="title"
                        value={titleValue}
                        onChange={handleTitleChange}
                        className="w-48 rounded-md border border-gray-300 p-2"
                    />
                </div>
                <div className="mb-2">
                    <label htmlFor="description">Description:</label>
                    <input
                        type="text"
                        id="description"
                        value={descriptionValue}
                        onChange={handleDescriptionChange}
                        className="h-48 w-96 rounded-md border border-gray-300 p-2"
                    />
                </div>
                <div className="mb-2">
                    <label htmlFor="ingredients">Ingredients:</label>
                    <input
                        type="text"
                        id="ingredients"
                        value={ingredientsValue}
                        onChange={handleIngredientsChange}
                        className="h-36 w-96 rounded-md border border-gray-300 p-2"
                    />
                </div>
            </div>
            <div>Picture:</div>
        </div>
        <div className="flex mb-2">
            <div className="mr-8">
                <label htmlFor="ingredients">Preparation time:</label>
                <input
                    type="text"
                    id="preparation-h"
                    className="w-12 rounded-md border border-gray-300 p-2"
                // value={preparationHValue}
                // onChange={handlePreparationHChange}
                />
                Hours
                <input
                    type="text"
                    id="preparation-h"
                    className="w-12 rounded-md border border-gray-300 p-2"
                // value={preparationHValue}
                // onChange={handlePreparationMChange}
                />
                Mins
            </div>
            <div>
                <label htmlFor="ingredients">Cooking time:</label>
                <input
                    type="text"
                    id="preparation-h"
                    className="w-12 rounded-md border border-gray-300 p-2"
                // value={cookingHValue}
                // onChange={handleCookingHChange}
                />
                Hours
                <input
                    type="text"
                    id="preparation-h"
                    className="w-12 rounded-md border border-gray-300 p-2"
                // value={cookingMValue}
                // onChange={handleCookingMChange}
                />
                Mins
            </div>
        </div>
        <div className="flex">
            Difficulty:
            <img src={StarImg} alt="star" />
        </div>
    </div>
}

export default NewRecipeModal;