import styles from './Step1BookName.module.css';

export type Props = {
  title: string;
  changeTitle: (event: React.ChangeEvent<HTMLInputElement>) => void;
  nextStep: () => void;
}

const Step1BookName = ({ title, changeTitle, nextStep }: Props) => {
  return (
    <div className={styles.container}>
      <h1 className={styles.header}>Let's think of a book title</h1>
      <input 
        id="title-input"
        type="text" 
        placeholder='Book Title'
        value={title}
        onChange={changeTitle}
        className={styles.input}
      />
      <button 
        onClick={nextStep}
        disabled={!title.trim()}
        className={styles.button}
      >
        Next
      </button>
     </div>
  );
}

export default Step1BookName;
