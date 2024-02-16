import leftArrow from './assets/left-arrow.svg';
import rightArrow from './assets/right-arrow.svg';

const Calendar = () => {

    const getDayWithSuffix = (day) => {
        if (day >= 11 && day <= 13) {
          return `${day}th`;
        }
      
        switch (day % 10) {
          case 1:
            return `${day}st`;
          case 2:
            return `${day}nd`;
          case 3:
            return `${day}rd`;
          default:
            return `${day}th`;
        }
      };

    const currentDate = new Date();

    // Get the current month
    const monthOptions = { month: 'long' };
    const formattedMonth = new Intl.DateTimeFormat('en-US', monthOptions).format(currentDate);

    const currentDay = currentDate.getDay(); // 0 for Sunday, 1 for Monday, ..., 6 for Saturday

    const difference = currentDay - 1 < 0 ? 6 : currentDay - 1;
    const startOfWeek = new Date(currentDate);
    startOfWeek.setDate(currentDate.getDate() - difference);

    const datesOfWeekWithSuffix = [];

    for (let i = 0; i < 7; i++) {
      const date = new Date(startOfWeek);
      date.setDate(startOfWeek.getDate() + i);
      const dayWithSuffix = getDayWithSuffix(date.getDate());
      datesOfWeekWithSuffix.push(date.toLocaleDateString('en-US', { weekday: 'long' }) + ' ' + dayWithSuffix);
    }
    console.log(datesOfWeekWithSuffix);

    return <div className="items-center">
        <div className="flex text-green-500">
            <img src={leftArrow} alt="left arrow mr-2" style={{fill: "red"}} />
            <div className="items-center">{formattedMonth}</div>
            <ul className="flex justify-between">{datesOfWeekWithSuffix.map((date, index) => <li key={index} className="border rounded-xl border-green-500 mx-2 px-2">{date}</li>)}</ul>
            <img src={rightArrow} alt="left arrow" />
        </div>
        <div>
            breakfast
            lunch
            dinner
        </div>
       
    </div>;
  };
  
  export default Calendar;