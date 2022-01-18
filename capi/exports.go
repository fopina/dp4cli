package capi

func ValidPWD(key1 []byte, key2 []byte, magicPin string) (string, int) {
	/*
	  int result; // eax

	  if ( !key1 )
	    return -5131;
	  if ( !key2 )
	    return -5132;
	  if ( !magicpin )
	    return -5133;
	  if ( !output )
	    return -5134;
	  result = sub_1D7760((int)key1, (int)key2, magicpin, (int)output);
	  if ( result >= 0 )
	    return sub_1D78A0(key1, (int *)key2, result);
	  return result;
	*/
	out, ret := myValidPWD_7760(key1, key2, magicPin)
	return out, ret
}
