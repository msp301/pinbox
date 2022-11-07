export interface Adapter<T> {
  adapt( item: any ): T;
}
