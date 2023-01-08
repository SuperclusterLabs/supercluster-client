import web3 from "web3";

export const truncate = (
  str: string | undefined,
  length: number
): string | undefined => {
  if (!str) {
    return str
  }
  if (str.length > length) {
    return `${str.substring(0, length - 3)}...`
  }
  return str
}

export const formatDate = (d: Date | undefined): string =>
  d ? d.toLocaleDateString('en-US') : ''

export const formatTime = (d: Date | undefined): string =>
  d
    ? d.toLocaleTimeString(undefined, {
      hour12: true,
      hour: 'numeric',
      minute: '2-digit',
    })
    : ''

export const checkIfPathIsEth = (address: string): boolean => {
  const { isAddress } = web3.utils;

  return isAddress(address);
}

export const checkIfPathIsEns = (address: string): boolean => {
  return address.includes('eth')
}

export const shortAddress = (addr: string): string =>
  addr.length > 10 && addr.startsWith('0x')
    ? `${addr.substring(0, 6)}...${addr.substring(addr.length - 4)}`
    : addr
