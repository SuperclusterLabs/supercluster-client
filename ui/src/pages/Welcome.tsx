import { MetamaskConnect } from "../components/MetamaskConnect";
import superclusterLogo from "../assets/superclusterLogo.svg";

function Welcome() {
  return (
    <div className="flex h-screen bg-onboarding-bg">
      <div className="m-auto text-center">
        <img
          className="max-w-none h-[37px]"
          src={superclusterLogo}
          alt="Supercluster logo"
        />
        <h1 className="text-7xl font-bold text-white mb-10">
          Supercluster Files
        </h1>
        <p className="text-2xl text-l-slategray-50">
          Share files with your team with maximum decentralization.
        </p>
        <MetamaskConnect />
      </div>
    </div>
  );
}

export default Welcome;