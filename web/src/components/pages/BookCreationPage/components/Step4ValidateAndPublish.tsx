export type Step4ValidateAndPublishProps = {
  title: string;
  previousStep: () => void;
  handlePublish: () => void;
};

const Step4ValidateAndPublish = ({ title, previousStep, handlePublish }: Step4ValidateAndPublishProps) => {
  return (
    <div>
      <h1>Validate your current story.</h1>
      <p>{title}</p>
      <button onClick={previousStep}>Prveious</button>
      <button onClick={handlePublish}>Publish</button>
     </div>
  );
}

export default Step4ValidateAndPublish;
