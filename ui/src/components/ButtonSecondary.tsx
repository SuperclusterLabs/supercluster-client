type Props = {
  text: string;
  onClick?:
    | ((event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void)
    | undefined;
};

function ButtonSecondary(props: Props) {
  return (
    <button
      className="bg-gradient-to-b from-l-slategray-200 to-l-slategray-300 py-4 px-14 rounded-2xl"
      onClick={props.onClick}
    >
      <span className="text-white font-bold text-md">{props.text}</span>
    </button>
  );
}

export default ButtonSecondary;
