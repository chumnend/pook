import styled from 'styled-components';
import { device, color } from '../../theme';

export const Page = styled.div`
  width: 100%;
  height: 100%;
  min-height: 768px;
  padding-top: 2rem;

  @media all and (min-width: ${device.lg}) {
    padding: 2rem;
  }
`;

export const StyledForm = styled.form`
  width: 100%;
  max-width: 520px;
  margin: 0 auto;
  padding: 2rem;
  border: none;
  & p {
    text-align: center;
  }

  @media all and (min-width: ${device.lg}) {
    width: 90%;
    border: 1px solid ${color.black};
  }
`;

export const StyledFormHeader = styled.div`
  margin-bottom: 1rem;
  text-align: center;
  & h2 {
    font-size: 1.5rem;
  }
  & p {
    margin: 1rem 0;
    padding: 0.8rem 0;
    background: ${color.red};
    color: ${color.black};
  }
`;

export const StyledFormGroup = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  margin-bottom: 1rem;
  & label {
    margin-bottom: 0.3rem;
    font-size: 1rem;
  }
  & input {
    padding: 0.8rem;
  }
  & small {
    font-size: 0.8rem;
    color: ${color.red};
  }
`;

export const StyledButton = styled.button`
  width: 100%;
  margin: 1rem auto;
  padding: 0.8rem 1rem;
  font-size: 1rem;
  cursor: pointer;
`;
