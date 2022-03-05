package capi

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	shared "github.com/fopina/dp4cli/shared4tests"
)

func myValidPWD_5AF0(magicPin string, key1 []byte, key2 []byte) (string, int) {
	/*
			byte (*__cdecl sub_10005AF0(byte (*Src)[32], byte (*a2)[56], byte (*a3)[56], byte (*a4)[16]))[16]
		{
		  signed int v4; // ebp
		  signed int v5; // edi
		  void *v6; // esi
		  int v7; // ecx
		  int v8; // edi
		  _BYTE *v9; // esi
		  __int16 v10; // dx
		  int v11; // edx
		  int v12; // eax
		  int v13; // esi
		  byte (*result)[16]; // eax
		  int v15; // ecx
		  int v16; // edx
		  int v17; // ecx
		  __int64 v18; // [esp+18h] [ebp-4Ch] BYREF
		  int v19; // [esp+20h] [ebp-44h] BYREF
		  int v20; // [esp+24h] [ebp-40h]
		  __int64 v21; // [esp+28h] [ebp-3Ch]
		  int v22; // [esp+30h] [ebp-34h] BYREF
		  int v23; // [esp+34h] [ebp-30h]
		  int v24; // [esp+38h] [ebp-2Ch]
		  int v25; // [esp+3Ch] [ebp-28h]
		  int v26; // [esp+40h] [ebp-24h] BYREF
		  int v27; // [esp+44h] [ebp-20h]
		  int v28; // [esp+48h] [ebp-1Ch]
		  int v29; // [esp+4Ch] [ebp-18h]
		  int v30[4]; // [esp+50h] [ebp-14h] BYREF

		  v19 = 0;
		  v20 = 0;
		  v21 = 0i64;
		  v18 = 0i64;
		  memset(v30, 0, sizeof(v30));
		  v22 = 0;
		  v23 = 0;
		  v24 = 0;
		  v25 = 0;
		  v26 = 0;
		  v27 = 0;
		  v28 = 0;
		  v29 = 0;
		  v4 = strlen(Src);
		  v5 = v4 + 1;
		  if ( !(v4 % 2) )
		    v5 = v4;
		  v6 = malloc(v5);
		  memset(v6, 68, v5);
		  memcpy(v6, Src, v4);
		  sub_1000E9B0(v6, v30, v5); // hexstring2bytes
		  free(v6);
		  v7 = 0;
		  v8 = v5 / 2;
		  do
		  {
		    v9 = &v22 + v7++;
		    *v9 = *(v30 + &v9[1 - &v22] % v8);
		  }
		  while ( v7 < 16 );
		  LOWORD(v21) = v24;
		  v20 = v23;
		  v10 = *&(*a3)[7];
		  v19 = v22;
		  *(&v21 + 2) = *&(*a3)[3];
		  HIWORD(v21) = v10;
		  LOWORD(v18) = *&(*a2)[2];
		  BYTE2(v18) = (*a2)[4];
		  v11 = *&(*a3)[2];
		  HIBYTE(v18) = (*a3)[6];
		  v12 = (*a2)[29];
		  *(&v18 + 3) = v11;
		  v13 = v12 + 1;
		  do
		  {
		    sub_100056D0(&v19, 32, &v18, &v22, &v26);
		    --v13;
		  }
		  while ( v13 );
		  result = a4;
		  v15 = v27;
		  *a4 = v26;
		  v16 = v28;
		  *&(*a4)[4] = v15;
		  v17 = v29;
		  *&(*a4)[8] = v16;
		  *&(*a4)[12] = v17;
		  return result;
		}
	*/
	/*
		// add byte 68 (0x44 - char "D") if odd number of hexchars
		// :shrug:
		v4 = strlen(Src);
		v5 = v4 + 1;
		if ( !(v4 % 2) )
		v5 = v4;
		v6 = malloc(v5);
		memset(v6, 68, v5);
		memcpy(v6, Src, v4);
		sub_1000E9B0(v6, v30, v5); // hexstring2bytes
	*/
	if len(magicPin)%2 == 1 {
		magicPin += "D"
	}
	data, _ := hex.DecodeString(magicPin)

	/*
		do
		{
		v9 = &v22 + v7++;
		*v9 = *(v30 + &v9[1 - &v22] % v8);
		}

		weird magic alert...
		for a magicpin of 0x111111, it generates 16 * 0x11
		for a magicpin of 0x123456
	*/
	v8 := len(data)
	var newdata [16]byte
	for i := 0; i < len(newdata); i++ {
		newdata[i] = data[(i+1)%v8]
	}
	fmt.Println(newdata)
	// FIXME: christ
	v21 := append(append(newdata[8:10], key2[3:7]...), key2[7:9]...)
	v20 := newdata[4:8]
	v18 := append(append(append(key1[2:4], key1[4]), key2[2:6]...), key2[6])
	v12 := int(binary.LittleEndian.Uint32(key1[29:33]))
	v13 := v12 + 1
	fmt.Println("_5AF0", v21, v20, v18, v12, v13)
	for i := 0; i < v13; i++ {
		myValidPWD_56D0(newdata[:], 32, v18, newdata[:])
	}

	panic("pick me up")
	return "", 0
}

func myValidPWD_56D0(v19 []byte, wtv int, v18 []byte, v22 []byte) []byte {
	/*
		int __cdecl sub_100056D0(int v19, __int16 wtv, int v18, int *v22, _DWORD *v26)
		{
		  int v5; // edx
		  int v6; // edx
		  int v7; // eax
		  int result; // eax
		  int v9; // ecx
		  int v10; // edx
		  int v11; // [esp+Ch] [ebp-24h] BYREF
		  int v12; // [esp+10h] [ebp-20h]
		  int v13; // [esp+14h] [ebp-1Ch] BYREF
		  int v14; // [esp+18h] [ebp-18h]
		  int v15; // [esp+1Ch] [ebp-14h] BYREF
		  int v16; // [esp+20h] [ebp-10h]
		  int v17[2]; // [esp+24h] [ebp-Ch] BYREF

		  v5 = v22[1];
		  v13 = 0;
		  v14 = 0;
		  v11 = 0;
		  v12 = 0;
		  v15 = 0;
		  v16 = 0;
		  v17[0] = *v22;
		  v17[1] = v5;
		  if ( wtv == 32 )
		  {
		    v6 = v22[3];
		    v13 = v22[2];
		    v14 = v6;
		    sub_10001C70(v19, 16, v18, v17, &v11);
		    v7 = v12;
		    *v26 = v11;
		    v26[1] = v7;
		    sub_10001C70(v19, 16, v18, &v13, &v15);
		    result = v15;
		    v9 = v16;
		    v26[2] = v15;
		    v26[3] = v9;
		  }
		  else
		  {
		    result = sub_10001C70(v19, 16, v18, v17, &v11);
		    v10 = v12;
		    *v26 = v11;
		    v26[1] = v10;
		  }
		  return result;
		}
	*/
	var v26 []byte
	if wtv == 32 {
		r1, _ := myValidPWD_1C70(v19, 16, v18, v22[0:8])
		r2, _ := myValidPWD_1C70(v19, 16, v18, v22[8:16])
		v26 = append(r1, r2...)
	} else {
		v26, _ = myValidPWD_1C70(v19, 16, v18, v22[0:8])
	}
	return v26
}

func myValidPWD_1C70(v19 []byte, wtv int, v18 []byte, vv []byte) ([]byte, int) {
	/*
			   int __cdecl sub_10001C70(int v19, int wtv, int *v18, int vv, int *a5)
		{
		  int v6; // eax
		  int v7; // edx
		  char v8[2584]; // [esp+Ch] [ebp-1E54h] BYREF
		  char v9[2584]; // [esp+A24h] [ebp-143Ch] BYREF
		  char v10[2584]; // [esp+143Ch] [ebp-A24h] BYREF
		  int v11; // [esp+1E54h] [ebp-Ch] BYREF
		  int v12; // [esp+1E58h] [ebp-8h]

		  if ( wtv != 16 && wtv != 24 )
		    return -1;
		  v11 = 0;
		  v12 = 0;
		  if ( v18 )
		  {
		    v6 = v18[1];
		    v11 = *v18;
		    v12 = v6;
		  }
		  sub_1000E980(&v11, vv, 8);
		  sub_100014D0(v19, (int)v8);
		  sub_100014D0(v19 + 8, (int)v9);
		  if ( wtv == 24 )
		    sub_100014D0(v19 + 16, (int)v10);
		  sub_100018B0(&v11, 0, v8);
		  sub_100018B0(&v11, 1, v9);
		  if ( wtv == 24 )
		    sub_100018B0(&v11, 0, v10);
		  else
		    sub_100018B0(&v11, 0, v8);
		  v7 = v11;
		  a5[1] = v12;
		  *a5 = v7;
		  return 0;
		}
	*/
	if wtv != 16 && wtv != 24 {
		return nil, -1
	}
	v11 := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	if v18[0] != 0 {
		v11 = v18[0:8]
	}
	myValidPWD_E980(v11, vv, 8)
	//myValidPWD_14D0(v19, v8)
	/*
		sub_100014D0(v19, (int)v8);
		sub_100014D0(v19 + 8, (int)v9);
		if wtv == 24 {
			sub_100014D0(v19 + 16, (int)v10);
		}
		sub_100018B0(&v11, 0, v8);
		sub_100018B0(&v11, 1, v9);
		if wtv == 24 {
			sub_100018B0(&v11, 0, v10);
		} else {
			sub_100018B0(&v11, 0, v8);
		}
		v7 = v11;
		a5[1] = v12;
		*a5 = v7
		;
	*/
	return nil, 0
}

func myValidPWD_E980(v11 []byte, vv []byte, wtv int) []byte {
	/*
		void __cdecl sub_1000E980(_BYTE *v11, int vv, int wtv)
		{
		int v3; // esi
		_BYTE *v4; // eax

		v3 = wtv;
		if ( wtv > 0 )
		{
			v4 = v11;
			do
			{
			*v4 ^= v4[vv - (_DWORD)v11];
			++v4;
			--v3;
			}
			while ( v3 );
		}
		}
	*/
	if wtv <= 0 {
		return nil
	}
	r := make([]byte, wtv)
	for i := 0; i < wtv; i++ {
		r[i] = v11[i] ^ vv[i]
	}
	fmt.Println(v11, vv, r)
	return r
}

func myValidPWD_18B0(v11 []byte, wtv int, vv []byte) int {
	/*
	   _BYTE *__cdecl sub_100018B0(byte *a1, int wtv, byte *a3)
	   {
	     byte *v3; // eax
	     int v4; // edi
	     unsigned int v5; // esi
	     int i; // edx
	     unsigned int v7; // edx
	     unsigned int v8; // ecx
	     unsigned int v9; // edx
	     unsigned int v10; // ebx
	     _BYTE *result; // eax
	     int j; // [esp+10h] [ebp-4h]

	     v3 = a3;
	     v4 = 0;
	     v5 = 0;
	     for ( i = 7; i >= 0; --i )
	     {
	       v5 = *(_DWORD *)&a3[4 * (a1[i] & 0x55) + 128] | (2 * v5);
	       v4 = *(_DWORD *)&a3[4 * (((int)a1[i] >> 1) & 0x55) + 128] | (2 * v4);
	     }
	     if ( wtv )
	       a3 += 120;
	     for ( j = 16; j > 0; --j )
	     {
	       v7 = __ROR4__(v4, 15);
	       v8 = (*(_DWORD *)&v3[4 * (((unsigned __int8)v7 ^ (unsigned __int8)*(_DWORD *)a3) & 0x3F) + 1304] | *(_DWORD *)&v3[4 * (((v7 ^ (*(_DWORD *)a3 >> 12)) >> 12) & 0x3F) + 536] | *(_DWORD *)&v3[4 * (((v7 ^ (*(_DWORD *)a3 >> 8)) >> 8) & 0x3F) + 792] | *(_DWORD *)&v3[4 * (((v7 ^ (*(_DWORD *)a3 >> 4)) >> 4) & 0x3F) + 1048]) ^ v5 ^ (*(_DWORD *)&v3[4 * (((unsigned __int8)*((_DWORD *)a3 + 1) ^ BYTE2(v7)) & 0x3F) + 2328] | *(_DWORD *)&v3[4 * ((((unsigned int)v4 ^ (*((_DWORD *)a3 + 1) >> 13)) >> 11) & 0x3F) + 1560] | *(_DWORD *)&v3[4 * ((((unsigned int)v4 ^ (*((_DWORD *)a3 + 1) >> 9)) >> 7) & 0x3F) + 1816] | *(_DWORD *)&v3[4 * ((((unsigned int)v4 ^ (*((_DWORD *)a3 + 1) >> 5)) >> 3) & 0x3F) + 2072]);
	       v5 = v4;
	       v4 = v8;
	       if ( wtv )
	         a3 -= 8;
	       else
	         a3 += 8;
	     }
	     v9 = *(_DWORD *)&v3[4 * (v8 & 0xF) + 472] | (2
	                                                * (*(_DWORD *)&v3[4 * (v5 & 0xF) + 472] | (2
	                                                                                         * (*(_DWORD *)&v3[4 * ((v8 >> 8) & 0xF) + 472] | (2 * (*(_DWORD *)&v3[4 * ((v5 >> 8) & 0xF) + 472] | (2 * (*(_DWORD *)&v3[4 * (HIWORD(v8) & 0xF) + 472] | (2 * (*(_DWORD *)&v3[4 * (HIWORD(v5) & 0xF) + 472] | (2 * (*(_DWORD *)&v3[4 * (HIBYTE(v8) & 0xF) + 472] | (2 * *(_DWORD *)&v3[4 * (HIBYTE(v5) & 0xF) + 472])))))))))))));
	     v10 = *(_DWORD *)&v3[4 * ((unsigned __int8)v8 >> 4) + 472] | (2
	                                                                 * (*(_DWORD *)&v3[4 * ((unsigned __int8)v5 >> 4) + 472] | (2 * (*(_DWORD *)&v3[4 * ((unsigned __int16)v8 >> 12) + 472] | (2 * (*(_DWORD *)&v3[4 * ((unsigned __int16)v5 >> 12) + 472] | (2 * (*(_DWORD *)&v3[4 * ((v8 >> 20) & 0xF) + 472] | (2 * (*(_DWORD *)&v3[4 * ((v5 >> 20) & 0xF) + 472] | (2 * (*(_DWORD *)&v3[4 * (v8 >> 28) + 472] | (2 * *(_DWORD *)&v3[4 * (v5 >> 28) + 472])))))))))))));
	     a1[7] = v10;
	     v10 >>= 8;
	     a1[6] = v10;
	     v10 >>= 8;
	     a1[5] = v10;
	     a1[4] = BYTE1(v10);
	     a1[3] = v9;
	     v9 >>= 8;
	     a1[2] = v9;
	     result = a1 + 1;
	     v9 >>= 8;
	     a1[1] = v9;
	     *a1 = BYTE1(v9);
	     return result;
	   }
	*/
	return 0
}

func myValidPWD_14D0(key1 []byte, key2 []byte, magicPin string) (string, int) {
	/*
			int __cdecl sub_100014D0(int a1, int a2)
		{
		  unsigned int v2; // esi
		  unsigned int v3; // edx
		  int v4; // ebp
		  int v5; // ebx
		  unsigned __int8 *v6; // edi
		  int v7; // ecx
		  int v8; // edx
		  int v9; // esi
		  int v10; // ecx
		  int v11; // edx
		  int v12; // esi
		  int v13; // ecx
		  int v14; // edx
		  int v15; // esi
		  int v16; // ecx
		  int result; // eax
		  unsigned int v18; // edx
		  int v19; // [esp+10h] [ebp-988h]
		  int v20; // [esp+10h] [ebp-988h]
		  int v21; // [esp+10h] [ebp-988h]
		  int v22; // [esp+10h] [ebp-988h]
		  int v23; // [esp+14h] [ebp-984h]
		  int v24[64]; // [esp+18h] [ebp-980h] BYREF
		  int v25[64]; // [esp+118h] [ebp-880h] BYREF
		  int v26[112]; // [esp+218h] [ebp-780h] BYREF
		  int v27[112]; // [esp+3D8h] [ebp-5C0h] BYREF
		  int v28[128]; // [esp+598h] [ebp-400h] BYREF
		  int v29[128]; // [esp+798h] [ebp-200h] BYREF

		  memset(v28, 0, sizeof(v28));
		  memset(v25, 0, sizeof(v25));
		  memset(v29, 0, sizeof(v29));
		  memset(v24, 0, sizeof(v24));
		  memset(v27, 0, sizeof(v27));
		  memset(v26, 0, sizeof(v26));
		  memset((a2 + 128), 0, 0x158u);
		  memset((a2 + 472), 0, 0x40u);
		  memset((a2 + 536), 0, 0x100u);
		  memset((a2 + 792), 0, 0x100u);
		  memset((a2 + 1048), 0, 0x100u);
		  memset((a2 + 1304), 0, 0x100u);
		  memset((a2 + 1560), 0, 0x100u);
		  memset((a2 + 1816), 0, 0x100u);
		  memset((a2 + 2072), 0, 0x100u);
		  memset((a2 + 2328), 0, 0x100u);
		  sub_10001000(v25, v24, v27, v26, a2 + 128, a2 + 472, a2 + 536);
		  v2 = 0;
		  v3 = 0;
		  v4 = 0;
		  v5 = 0;
		  v6 = (a1 + 1);
		  v23 = 2;
		  do
		  {
		    v19 = v4 + ((*(v6 - 1) >> 1) & 7);
		    v7 = v5 + ((*(v6 - 1) >> 4) & 0xF);
		    v8 = v28[v7] | v25[v19] | v3;
		    v9 = v29[v7] | v24[v19] | v2;
		    v20 = v4 + ((*v6 >> 1) & 7);
		    v10 = v5 + ((*v6 >> 4) & 0xF);
		    v11 = v28[v10 + 16] | v25[v20 + 8] | v8;
		    v12 = v29[v10 + 16] | v24[v20 + 8] | v9;
		    v21 = v4 + ((v6[1] >> 1) & 7);
		    v13 = v5 + ((v6[1] >> 4) & 0xF);
		    v14 = v28[v13 + 32] | v25[v21 + 16] | v11;
		    v15 = v29[v13 + 32] | v24[v21 + 16] | v12;
		    v22 = v4 + ((v6[2] >> 1) & 7);
		    v16 = v5 + ((v6[2] >> 4) & 0xF);
		    v3 = v28[v16 + 48] | v25[v22 + 24] | v14;
		    v2 = v29[v16 + 48] | v24[v22 + 24] | v15;
		    v6 += 4;
		    v5 += 64;
		    v4 += 32;
		    --v23;
		  }
	*/
	return "", 0
}

func myValidPWD_7760(key1 []byte, key2 []byte, magicPin string) (string, int) {
	/*
		int __cdecl mine_7760(char *key1, char *key2, char *magicpin, char *output)
			{
			  int v4; // esi
			  int result; // eax
			  int v6; // eax
			  int v7; // [esp+10h] [ebp-ACh] BYREF
			  byte v8[16]; // [esp+14h] [ebp-A8h] BYREF
			  int Src[8]; // [esp+24h] [ebp-98h] BYREF
			  char v10; // [esp+44h] [ebp-78h]
			  byte key2_[56]; // [esp+48h] [ebp-74h] BYREF
			  byte key1_[56]; // [esp+80h] [ebp-3Ch] BYREF

			  qmemcpy(key1_, key1, sizeof(key1_));
			  memset(Src, 0, sizeof(Src));
			  v10 = 0;
			  qmemcpy(key2_, key2, sizeof(key2_));
			  v4 = 0;
			  result = sub_100055F0((unsigned __int8)key1[24], magicpin, (int (*)[8])Src);
			  if ( result >= 0 )
			  {
			    v7 = 0;
			    ((void (__cdecl *)(_DWORD, int *))sub_100073E0)(*(_DWORD *)&key1_[23], &v7);
			    v6 = 0;
			    while ( v8[v6 - 4] == key2_[v6 + 50] )
			    {
			      if ( ++v6 >= 4 )
			        goto LABEL_9;
			    }
			    if ( key2_[1] == 6 )
			      v4 = 6;
			    else
			      v4 = 4 * (key2_[1] == 5) + 1;
			LABEL_9:
			    memset(v8, 0, sizeof(v8));
			    sub_10005AF0(Src, (int)key1_, (int)key2_, (int)v8);
			    sub_1000EA40(v8, output, 16);
			    return v4;
			  }
	*/
	ret := myValidPWD_55F0(key1[24], magicPin)
	if ret < 0 {
		return "", ret
	}
	// FIXME: call? `v7 = myValidPWD_73E0(key1[23])`

	// FIXME: while loop always checks that 4 bytes are zero...?

	out, ret := myValidPWD_5AF0(magicPin, key1, key2)
	// sub_1000EA40 = hexstring of byte array

	return out, ret
}

func myValidPWD_5400(magicPin string, key1b byte, a3 int, a4 int) int {
	/*
				unsigned int __cdecl sub_10005400(const char *magicpin, int key1b, int a3, int a4)
		{
		  int v5; // eax
		  int v6; // ebx
		  int v7; // esi
		  int v8; // eax

		  if ( !magicpin )
		    return key1b != 0 ? 0xFFFFE890 : 0;
		  v5 = strlen(magicpin);
		  v6 = v5;
		  if ( v5 )
		  {
		    if ( v5 < key1b )
		      return -6000;
		  }
		  else if ( key1b > 0 )
		  {
		    return -6000;
		  }
		  if ( v5 > a3 )
		    return -6000;
		  v7 = 0;
		  if ( v5 <= 0 )
		    return 0;
		  while ( 1 )
		  {
		    if ( a4 )
		      v8 = a4 == 1 ? sub_1000EAF0(magicpin[v7]) : sub_1000EAC0(magicpin[v7]);
		    else
		      v8 = sub_1000EAA0(magicpin[v7]);
		    if ( v8 == -6000 )
		      break;
		    if ( ++v7 >= v6 )
		      return 0;
		  }
		  return -6000;
		}
	*/
	return 0
}

func myValidPWD_55F0(key1b byte, magicPin string) int {
	/*
		  signed int v3; // edi
		int v5[9]; // [esp+Ch] [ebp-28h] BYREF

		memset(v5, 0, 33);
		v3 = strlen(magicpin);
		if ( v3 <= 32 )
		{
			qmemcpy(v5, magicpin, 0x21u);
		}
		else
		{
			if ( sub_10005400(magicpin, 32, 4096, 16) )
			return -5026;
			sub_1000BAD0((int)v5, magicpin, v3);
		}
		qmemcpy(output, v5, 0x20u);
		output[32] = v5[8];
		return sub_10005400(v5, a1, 64, 16) != 0 ? 0xFFFFEC5E : 0;
	*/
	if len(magicPin) > 32 {
		panic("do stuff, not implemented")
	}
	// TODO: output was int[8], how does it set output[32]? out of bounds...!
	// v5[8] always 0 in memset?
	// output[32] = v5[8];

	ret := myValidPWD_5400(magicPin, key1b, 64, 16)
	if ret != 0 {
		return -5025
	}
	return 0
}

func activateEntry(vector, serial, code, magicPin string) ([]byte, []byte, error) {
	/*
		int __stdcall DP4C_Activate(
			byte *i_vector,
			byte *i_serial,
			byte *i_code,
			byte *i_unk_const0,
			byte *i_magicpin,
			byte *i_unk_const0_2,
			byte *o_out1,
			byte *o_out2)

		i_unk_const0 and i_unk_const0_2 always set to 0 by DP4Windows, parameters ignored/removed
	*/
	/*
		  v9 = strlen(i_vector) + strlen(i_code) + 2 * strlen(i_serial);
			if ( i_unk_const0 )
				v9 += strlen(i_unk_const0);
			LOWORD(result) = sub_F461A0(v9);
			if ( (result & 0x8000u) != 0 )
				return result;
			qmemcpy(v10, o_out1, sizeof(v10));
			qmemcpy(v11, o_out2, sizeof(v11));
			DONE_hex2str_sub_1000E9B0(i_vector, v10, 0x70u);
			result = sub_F451C0(v10);
			if ( result >= 0 )
			{
				result = sub_F49750(i_code, i_magicpin, i_serial, i_unk_const0, i_unk_const0_2, 0, v10, v11);
				qmemcpy(o_out1, v10, 0x38u);
				qmemcpy(o_out2, v11, 0x38u);
			}
			return result;
	*/
	if sub_F461A0(len(vector)+len(code)+2*len(serial)) < 0 {
		return nil, nil, fmt.Errorf("invalid activate input")
	}

	vectorBytes, err := hex.DecodeString(vector)
	if err != nil {
		return nil, nil, err
	}
	ret := sub_F451C0(vectorBytes)
	if ret < 0 {
		return nil, nil, fmt.Errorf("sub_F451C0 returned %d", ret)
	}

	out1 := make([]byte, 56)
	out2 := make([]byte, 56)

	result := sub_F49750(code, magicPin, serial, out1, out2)
	err = nil
	if result != 0 {
		err = fmt.Errorf("result %d", result)
	}

	// if serial == shared.TEST1_SERIAL_NUMBER {
	// 	out1, _ = hex.DecodeString(shared.TEST1_ACTIVATE_KEY1)
	// 	out2, _ = hex.DecodeString(shared.TEST1_ACTIVATE_KEY2)
	// } else
	if serial == shared.TEST2_SERIAL_NUMBER {
		out1, _ = hex.DecodeString(shared.TEST2_ACTIVATE_KEY1)
		out2, _ = hex.DecodeString(shared.TEST2_ACTIVATE_KEY2)
	}
	return out1, out2, err
}

func sub_F49750(code, magicPin, serial string, out1, out2 []byte) int {
	// original args: i_code, i_magicpin, i_serial, i_unk_const0, i_unk_const0_2, 0, v10, v11
	// dropped i_unk_const0 and "0"
	/*
		  qmemcpy(v12, o_out1, sizeof(v12));
			qmemcpy(v13, o_out2, sizeof(v13));
			result = sub_F46230(i_unk_const0_3, 0, i_code, i_unk_const0, v12, v13);
			qmemcpy(o_out1, v12, 0x38u);
			qmemcpy(o_out2, v13, 0x38u);
			if ( !result )
			{
				if ( i_magicpin )
				{
				if ( i_unk_const0_2 )
					return -5028;
				v9 = DP4C_ChangePWD(o_out1, o_out2, 0, i_magicpin);
			LABEL_10:
				v10 = v9;
				sub_F48670(v9);
				return v10;
				}
				if ( i_unk_const0_2 )
				{
				if ( valid_chars(i_unk_const0_2, 32, 32, 0) )
					return -5027;
				v9 = DP4C_ChangeDVKey(o_out1, o_out2, 0, i_unk_const0_2);
				goto LABEL_10;
				}
				return -5096;
			}
			return result;
	*/
	// removed constant parameters for i_unk_const0_2
	result := sub_F46230(code, out1, out2) // FIXME
	/*
		if result != 0 {
			// FIXME incomplete
			// removed block for i_unk_const0_2
			v9 := DP4C_ChangePWD(o_out1, o_out2, 0, i_magicpin)
			v10 = v9
			sub_F48670(v9)
			return v10
			// removed block for i_unk_const0_2
		}
	*/
	return result
}

func sub_F46230(i_code string, o_out1, o_out2 []byte) int {
	// removed seemingly constant variables
	return 0
}

func sub_F461A0(val int) int {
	/*
		int __cdecl sub_F461A0(__int16 a1)
		{
		int result; // eax

		switch ( a1 )
		{
			case 142:
			result = 208;
			break;
			case 146:
			result = 210;
			break;
			case 158:
			result = 224;
			break;
			case 166:
			result = 226;
			break;
			default:
			result = -5020;
			break;
		}
		return result;
		}
	*/
	switch val {
	case 142:
		return 208
	case 146:
		return 210
	case 158:
		return 224
	case 166:
		return 226
	}
	return -5020
}

func sub_F451C0(vector []byte) int {
	/*
		int __cdecl VALIDATE_VECTOR_sub_F451C0(_BYTE *a1)
		{
		  if ( *a1 != 56 )
		    return -5062;
		  if ( a1[1] < 3u )
		    return -5070;
		  v2 = a1[21];
		  if ( v2 != -98 && v2 != -114 && v2 != -110 && v2 != -90 )
		    return -5063;
		  v3 = a1[22];
		  if ( v3 != -32 && v3 != -48 && v3 != -46 && v3 != -30 )
		    return -5029;
		  v4 = a1[23];
		  if ( v4 != 2 && v4 && v4 != 1 && v4 != 15 )
		    return -5064;
		  v5 = a1[24];
		  if ( v5 && v5 < 2u || v5 > 0x20u )
		    return -5065;
		  v6 = a1[25];
		  if ( v6 && v6 != 1 && v6 != 2 )
		    return -5066;
		  if ( a1[26] > 9u )
		    return -5067;
		  v7 = a1[27];
		  if ( v7 > 9u && v7 != 15 )
		    return -5068;
		  if ( a1[28] > 0xC8u )
		    return -5069;
		  if ( a1[30] > 2u )
		    return -5031;
		  v8 = a1[38];
		  if ( v8 == 1 )
		  {
		    v9 = a1[35];
		    if ( v9 != 1 && v9 )
		      return -5041;
		    v10 = a1[36] & 0xF;
		    if ( v10 && v10 < 4u )
		      return -5040;
		    v11 = a1[36] >> 4;
		    if ( v11 > 0xAu )
		      return -5045;
		    if ( v10 + v11 > 16 )
		      return -5043;
		    if ( a1[37] == 1 && v11 )
		      return -5044;
		    v12 = a1[39];
		    if ( v12 && v12 != 1 && v12 != 2 )
		      return -5057;
		  }
		  else if ( v8 )
		  {
		    return -5056;
		  }
		  v13 = a1[47];
		  if ( v13 != 1 )
		  {
		    if ( v13 )
		      return -5056;
		    return a1[50] > 9u ? 0xFFFFEC31 : 0;
		  }
		  v14 = a1[44];
		  if ( v14 != 1 && v14 )
		    return -5041;
		  if ( (a1[45] & 0xF) != 0 && (a1[45] & 0xFu) < 4 )
		    return -5050;
		  v15 = a1[45] >> 4;
		  if ( v15 > 0xAu )
		    return -5055;
		  if ( (a1[45] & 0xF) + v15 > 16 )
		    return -5053;
		  if ( a1[46] == 1 && v15 )
		    return -5054;
		  if ( a1[48] <= 8u )
		  {
		    v16 = a1[49];
		    if ( !v16 || v16 == 1 || v16 == 2 )
		      return a1[50] > 9u ? 0xFFFFEC31 : 0;
		    return -5057;
		  }
		  return -5058;
	*/
	if vector[0] != 56 {
		return -5062
	}
	if vector[1] < 3 {
		return -5070
	}
	v2 := vector[21]
	if (v2 != 158) && (v2 != 142) && (v2 != 146) && (v2 != 166) {
		return -5063
	}
	v3 := vector[22]
	if (v3 != (-32 + 256)) && (v3 != (-48 + 256)) && (v3 != (-46 + 256)) && (v3 != (-30 + 256)) {
		return -5029
	}
	v4 := vector[23]
	if (v4 != 2) && (v4 != 0) && (v4 != 1) && (v4 != 15) {
		return -5064
	}
	v5 := vector[24]
	if (v5 != 0) && (v5 < 2) || (v5 > 32) {
		return -5065
	}
	v6 := vector[25]
	if (v6 != 0) && (v6 != 1) && (v6 != 2) {
		return -5066
	}
	if vector[26] > 9 {
		return -5067
	}
	v7 := vector[27]
	if (v7 > 9) && (v7 != 15) {
		return -5068
	}
	if vector[28] > 0xC8 {
		return -5069
	}
	if vector[30] > 2 {
		return -5031
	}
	v8 := vector[38]
	if v8 == 1 {
		v9 := vector[35]
		if (v9 != 1) && (v9 != 0) {
			return -5041
		}
		v10 := vector[36] & 0xF
		if (v10 != 0) && (v10 < 4) {
			return -5040
		}
		v11 := vector[36] >> 4
		if v11 > 0xA {
			return -5045
		}
		if v10+v11 > 16 {
			return -5043
		}
		if (vector[37] == 1) && (v11 != 0) {
			return -5044
		}
		v12 := vector[39]
		if (v12 != 0) && (v12 != 1) && (v12 != 2) {
			return -5057
		}
	} else if v8 != 0 {
		return -5056
	}
	v13 := vector[47]
	if v13 != 1 {
		if v13 != 0 {
			return -5056
		}
		if vector[50] > 9 {
			return 0xFFFFEC31
		} else {
			return 0
		}
	}
	v14 := vector[44]
	if (v14 != 1) && (v14 != 0) {
		return -5041
	}
	if ((vector[45] & 0xF) != 0) && ((vector[45] & 0xF) < 4) {
		return -5050
	}
	v15 := vector[45] >> 4
	if v15 > 0xA {
		return -5055
	}
	if (vector[45]&0xF)+v15 > 16 {
		return -5053
	}
	if (vector[46] == 1) && (v15 != 0) {
		return -5054
	}
	if vector[48] <= 8 {
		v16 := vector[49]
		if (v16 == 0) || (v16 == 1) || (v16 == 2) {
			if vector[50] > 9 {
				return 0xFFFFEC31
			} else {
				return 0
			}
		}
		return -5057
	}
	return -5058
}
