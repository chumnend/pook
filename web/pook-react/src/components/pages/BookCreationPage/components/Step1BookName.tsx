export type Props = {
  title: string;
  changeTitle: (event: React.ChangeEvent<HTMLInputElement>) => void;
  nextStep: () => void;
}

const Step1BookName = ({ title, changeTitle, nextStep }: Props) => {
  return (
    <div>
      <h1>Let's think of a book title.</h1>
      <input 
        id="title-input"
        type="text" 
        placeholder='Book Title'
        value={title}
        onChange={changeTitle}
      />
      <button onClick={nextStep}>Next</button>
     </div>
  );
}

export default Step1BookName;
