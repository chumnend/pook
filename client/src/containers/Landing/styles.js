import styled from 'styled-components';

export const Page = styled.div`
  width: 100%;
  min-height: 768px;
`;

export const HeroImage = styled.div`
  width: 100%;
  height: 768px;
  background-image: linear-gradient(rgba(0, 0, 0, 0.3), rgba(0, 0, 0, 0.3)),
    url('https://i.imgur.com/0kKrFNV.jpg');
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  position: relative;
`;

export const HeroText = styled.div`
  text-align: center;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: white;
`;
